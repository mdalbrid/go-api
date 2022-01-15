package dto

import (
	"encoding/json"

	"github.com/mdalbrid/utils/server"
	"github.com/mdalbrid/utils/validators"
)

func ParamsToDto(params interface{}, requestDTO interface{}) (e error) {
	var paramsJson []byte
	paramsJson, e = json.Marshal(params)
	if e != nil {
		return
	}
	e = json.Unmarshal(paramsJson, &requestDTO)
	return
}

func ParseRequestParams(params interface{}, requestDTO interface{}) *server.JsonRpcError {
	err := ParamsToDto(params, &requestDTO)
	if err != nil {
		return &server.JsonRpcInvalidParamsError
	}
	err = validators.ValidateStruct(requestDTO)
	if err != nil {
		return &server.JsonRpcError{
			Code:    server.JsonRpcInvalidParamsError.Code,
			Message: server.JsonRpcInvalidParamsError.Message,
			Data:    err.Error(),
		}
	}
	return nil
}
