package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	logger "github.com/cihub/seelog"
)

// @Description http的POST请求
// @Author zhouhaiping
// @Version 1.0
// @Update 2020/5/21 21:36 init
// @Update Ryen 2020/5/23 15:33 增加超时控制, 规范命名
func HttpPost(url string, reqBody []byte, timeout time.Duration) ([]byte, error) {
	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		logger.Error("new http request error: ", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("http request error: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Get http response body error: ", err)
		return nil, err
	}

	return body, nil
}

// @Description http的GET请求
// @Author zhouhaiping
// @Version 1.0
// @Update 2020/5/21 21:37 init
// @Update Ryen 2020/5/23 15:33 增加超时控制, 规范命名
func HttpGet(url string, timeout time.Duration) (string, error) {
	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.Error("new http request error: ", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("http request error: ", err)
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Get http response body error: ", err)
		return "", err
	}

	return string(body), nil
}
