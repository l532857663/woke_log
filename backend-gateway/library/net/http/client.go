package http

import (
	"backend-gateway/utils"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	logger "github.com/cihub/seelog"
	"go.elastic.co/apm/module/apmhttp"
)

type Request struct {
	BaseUrl      string
	Headers      map[string]string
	HttpClient   *http.Client
	ErrorHandler func(res *http.Response, uri string) error
	ResponseTime int64
}

type HttpClientParam struct {
	BaseUrl     string
	RpcUser     string
	RpcPwd      string
	IsBasicAuth bool
	ApiKey      string
}

func (r *Request) GetResponseTime(startTime time.Time) {
	start := startTime.UnixNano() / 1e6
	end := time.Now().UnixNano() / 1e6
	r.ResponseTime = end - start
	return
}

func (r *Request) SetTimeout(seconds time.Duration) {
	r.HttpClient.Timeout = time.Second * seconds
}

func (r *Request) SafeDecodeResponse(body io.ReadCloser, res interface{}) (err error) {
	var data []byte
	defer func() {
		if r := recover(); r != nil {
			logger.Error("unmarshal json recovered from panic: ", r, "; data: ", string(data))
			if len(data) > 0 && len(data) < 2048 {
				err = fmt.Errorf("Error: %v", string(data))
			} else {
				err = fmt.Errorf("Internal error")
			}
		}
	}()
	data, err = ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &res)
}

func InitHttpClient4Blockbook(blockbookApiUrl string) Request {
	return Request{
		Headers:      make(map[string]string),
		HttpClient:   DefaultClient,
		ErrorHandler: DefaultErrorHandler,
		BaseUrl:      blockbookApiUrl,
	}
}

func InitHttpClientWithJSON(baseUrl, rpcUser, rpcPwd string, isBasicAuth bool) Request {
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	if isBasicAuth {
		headers["Authorization"] = "Basic " + basicAuth(rpcUser, rpcPwd)
	}

	return Request{
		Headers:      headers,
		HttpClient:   DefaultClient,
		ErrorHandler: DefaultErrorHandler,
		BaseUrl:      baseUrl,
	}
}

func InitHttpClientForOkex(param *HttpClientParam) Request {
	request := InitHttpClientWithJSON(param.BaseUrl, param.RpcUser, param.RpcPwd, param.IsBasicAuth)
	// request.Headers["x-apiKey"] = param.ApiKey

	return request
}

var DefaultClient = &http.Client{
	Timeout: time.Second * 5,
}

// See 2 (end of page 4) https://www.ietf.org/rfc/rfc2617.txt
// "To receive authorization, the client sends the userid and password,
// separated by a single colon (":") character, within a base64
// encoded string in the credentials."
// It is not meant to be urlencoded.
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

/*
func defaultBlockbookClient(crt, key string, insecureSkipVerify bool) *http.Client {
	cert, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		panic("Unable to load blockbook .crt and .key certs")
	}

	certBytes, err := ioutil.ReadFile(crt)
	if err != nil {
		panic("Unable to read blockbook.crt")
	}

	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		panic("failed to parse root certificate")
	}

	client := &http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            clientCertPool,
				Certificates:       []tls.Certificate{cert},
				InsecureSkipVerify: insecureSkipVerify,
			},
		},
	}

	return client
}
*/

var DefaultErrorHandler = func(res *http.Response, uri string) error {
	return nil
}

func (r *Request) GetWithContext(result interface{}, path string, query url.Values, ctx context.Context) error {
	var queryStr = ""
	if query != nil {
		queryStr = query.Encode()
	}

	uri := strings.Join([]string{r.GetBase(path), queryStr}, "?")
	return r.Execute("GET", uri, nil, result, ctx)
}

func (r *Request) Get(result interface{}, path string, query url.Values) error {
	var queryStr = ""
	if query != nil {
		queryStr = query.Encode()
	}

	uri := strings.Join([]string{r.GetBase(path), queryStr}, "?")
	return r.Execute("GET", uri, nil, result, context.Background())
}

func (r *Request) Post(result interface{}, path string, body interface{}) error {
	buf, err := GetBody(body)
	if err != nil {
		return err
	}

	uri := r.GetBase(path)
	return r.Execute("POST", uri, buf, result, context.Background())
}

func (r *Request) PostWithContext(result interface{}, path string, body interface{}, ctx context.Context) error {
	buf, err := GetBody(body)
	if err != nil {
		return err
	}

	uri := r.GetBase(path)
	return r.Execute("POST", uri, buf, result, ctx)
}

func (r *Request) Execute(method string, url string, body io.Reader, result interface{}, ctx context.Context) error {
	logger.Debugf("Execute uri: %v", url)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		logger.Errorf("Platform request error: %+v", err)
		return utils.ErrTypePlatformRequest(err)
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	c := apmhttp.WrapClient(r.HttpClient)

	start := time.Now()
	res, err := c.Do(req.WithContext(ctx))
	if err != nil {
		return utils.ErrTypePlatformRequest(err)
	}
	defer res.Body.Close()
	r.GetResponseTime(start)

	err = r.ErrorHandler(res, url)
	if err != nil {
		return utils.ErrTypePlatformError
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return utils.ErrTypePlatformUnmarshal
	}

	err = json.Unmarshal(b, result)
	if err != nil {
		logger.Errorf("Body unmarshal error: %+v, response body: %s", err, string(b))
		return utils.ErrTypePlatformUnmarshal
	}

	return err
}

func (r *Request) GetBase(path string) string {
	if path == "" {
		return r.BaseUrl
	}
	return fmt.Sprintf("%s/%s", r.BaseUrl, path)
}

func GetBody(body interface{}) (buf io.ReadWriter, err error) {
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
	}
	return
}
