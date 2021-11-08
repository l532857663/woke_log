package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	logger "github.com/cihub/seelog"

	"github.com/gin-gonic/gin"
)

type HttpLogFilterType string // 忽略类型

const (
	HTTP_LOG_FILTER_NULL HttpLogFilterType = "0" // 不忽略
	HTTP_LOG_FILTER_REQ  HttpLogFilterType = "1" // 忽略请求
	HTTP_LOG_FILTER_RESP HttpLogFilterType = "2" // 忽略响应
	HTTP_LOG_FILTER_ALL  HttpLogFilterType = "3" // 忽略所有

)

type HttpLogFilter struct {
	InterfacePath string            // 接口路径
	InterfaceDesc string            // 接口描述
	FilterType    HttpLogFilterType // 过滤器类型
}

var httpLoginFilterList = []HttpLogFilter{
	HttpLogFilter{
		InterfacePath: "/api/bingoo/wallet/accounts/v3/address",
		InterfaceDesc: "钱包三期首页列表",
		FilterType:    HTTP_LOG_FILTER_RESP,
	},
}

type MiddlewareFunc func(*gin.Context) error

// 中间件拦截请求验证Token，验证不通过不进入业务层
func TokenHandler(mf MiddlewareFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if mf != nil {
			err := mf(c)
			if err != nil {
				c.AbortWithStatus(http.StatusNetworkAuthenticationRequired)
				return
			}
		}

		// 处理请求
		c.Next()
	}
}

// 跨域资源共享
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-Auth-Token, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
		return
	}
}

// responseLogger 用于存放应答内容
type responseLogger struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (rl *responseLogger) Write(b []byte) (int, error) {
	rl.body.Write(b)
	return rl.ResponseWriter.Write(b)
}

func (rl *responseLogger) WriteString(s string) (int, error) {
	rl.body.WriteString(s)
	return rl.ResponseWriter.WriteString(s)
}

// 检查 API 请求 HEADER 和 BODY 信息，方便调试
func LogHttpWrapper() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		var err error
		logInfo := map[string]interface{}{
			"\n[METHOD]": c.Request.URL.String(),
			"\n[TIME]":   t.String(),
		}

		// 查询忽略列表
		filterFlag := HTTP_LOG_FILTER_NULL
		for _, value := range httpLoginFilterList {
			if strings.Contains(c.Request.URL.String(), value.InterfacePath) == true {
				filterFlag = value.FilterType
			}
		}

		if filterFlag == HTTP_LOG_FILTER_REQ || filterFlag == HTTP_LOG_FILTER_ALL {
			req, err := httputil.DumpRequest(c.Request, false)
			if err != nil {
				logInfo["\n[ERROR]"] = err
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			logInfo["\n[REQUEST]"] = string(req) + "BODY: ignore"
		} else {
			req, err := httputil.DumpRequest(c.Request, true)
			if err != nil {
				logInfo["\n[ERROR]"] = err
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			logInfo["\n[REQUEST]"] = string(req)
		}

		rlw := &responseLogger{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = rlw

		defer func() {
			// 格式化输出 方便查看
			var out bytes.Buffer
			err = json.Indent(&out, rlw.body.Bytes(), "", "    ")
			if err != nil {
				logger.Errorf("json indent error: %+v", err)
				return
			}
			// logger.Debugf("response body: %s", rlw.body.String())

			logInfo["\n[COST_TIME]"] = time.Now().Sub(t).String()

			if filterFlag == HTTP_LOG_FILTER_RESP || filterFlag == HTTP_LOG_FILTER_ALL {
				logInfo["\n[RESPONSE]"] = "res had ignored"
			} else {
				logInfo["\n[RESPONSE]"] = out.String()
			}

			if err != nil {
				logger.Error(err)
			} else {
				logger.Debug(logInfo)
			}
		}()

		c.Next()

	}
}
