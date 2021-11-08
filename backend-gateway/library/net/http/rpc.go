package http

import (
	"backend-gateway/utils"
	"encoding/json"
	"strconv"
	"strings"

	logger "github.com/cihub/seelog"
)

var (
	requestId = int64(0)
)

const (
	JSON_RPC_VERSION = "2.0"
)

type (
	RpcRequests []*RpcRequest

	RpcRequest struct {
		JsonRpc string      `json:"jsonrpc"`
		Method  string      `json:"method"`
		Params  interface{} `json:"params,omitempty"`
		Id      int64       `json:"id,omitempty"`
	}

	RpcResponse struct {
		JsonRpc string      `json:"jsonrpc"`
		Error   *RpcError   `json:"error,omitempty"`
		Result  interface{} `json:"result,omitempty"`
		Id      int64       `json:"id,omitempty"`
	}

	RpcError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func (r *Request) RpcCall(result interface{}, method string, params interface{}) error {
	req := &RpcRequest{JsonRpc: JSON_RPC_VERSION, Method: method, Params: params, Id: genId()}
	var resp *RpcResponse
	err := r.Post(&resp, "", req)
	if err != nil {
		return err
	}

	if resp.Error != nil {
		return utils.ErrRPCCallError(resp.Error.Code, resp.Error.Message)
	}

	return resp.GetObject(result)
}

func (r *RpcResponse) GetObject(toType interface{}) error {
	js, err := json.Marshal(r.Result)
	if err != nil {
		return utils.ErrRPCGetObjectMarshalError(err)
	}

	err = json.Unmarshal(js, toType)
	if err != nil {
		return utils.ErrRPCGetObjectUnMarshalError(err)
	}

	return nil
}

func genId() int64 {
	requestId += 1
	return requestId
}

func ParseRpcErrorCode(err error) (int64, error) {
	rpcErrorMessage := err.Error()
	errorInfosnippet := strings.Split(rpcErrorMessage, ",")
	for _, m := range errorInfosnippet {
		if strings.Contains(m, "code") {
			// 截取 code: 后的内容并返回
			codeStr := m[strings.IndexByte(m, '-'):]
			code, err := strconv.ParseInt(codeStr, 10, 64)
			if err != nil {
				logger.Error("Parse rpc error code error: ", err)
				return 0, err
			}
			return code, nil
		}
	}

	return 0, nil
}
