package main

import (
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/mdalbrid/go-api-aws/dto"
	objectGroupAttributeDto "github.com/mdalbrid/go-api-aws/dto/modules/object_attribute"
	"github.com/mdalbrid/models/db"
	"github.com/mdalbrid/models/ws_object_attribute"
	"github.com/mdalbrid/utils/logger"
	"github.com/mdalbrid/utils/server"
	"net/http"
)

func listHandler(req *server.JsonRpcRequest, res *server.JsonRpcResponse) int {
	listDTO := &objectGroupAttributeDto.ListRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, listDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	list, err := ws_object_attribute.List(listDTO.ToFilter())
	if err != nil {
		logger.Error(err)
		res.Error = &server.JsonRpcInternalError
		return http.StatusInternalServerError
	}

	total, err := ws_object_attribute.Count(listDTO.ToFilter())
	if err != nil {
		logger.Error(err)
		res.Error = &server.JsonRpcInternalError
		return http.StatusInternalServerError
	}

	res.Result = map[string]interface{}{
		"list":  list,
		"total": total,
	}

	return http.StatusOK
}

func createHandler(req *server.JsonRpcRequest, res *server.JsonRpcResponse) int {
	createDTO := &objectGroupAttributeDto.CreateRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, createDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	exp, err := ws_object_attribute.Create(createDTO.ObjectUUID, createDTO.Name, createDTO.Value)
	if err != nil {
		logger.Error(err)
		res.Error = &server.JsonRpcInternalError
		return http.StatusInternalServerError
	}

	res.Result = exp
	return http.StatusOK
}

func editHandler(req *server.JsonRpcRequest, res *server.JsonRpcResponse) int {
	editDTO := &objectGroupAttributeDto.EditRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, editDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	exp, err := ws_object_attribute.Edit(editDTO.Uuid, editDTO.Name, editDTO.Value, editDTO.SortWeight)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			res.Error = &server.JsonRpcInvalidParamsError
			res.Error.Data = db.ErrNotFound
			return http.StatusNotFound
		}
		logger.Error(err)
		res.Error = &server.JsonRpcInternalError
		return http.StatusInternalServerError
	}

	res.Result = exp
	return http.StatusOK
}

func deleteHandler(req *server.JsonRpcRequest, res *server.JsonRpcResponse) int {
	deleteDTO := &objectGroupAttributeDto.DeleteRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, deleteDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	exp, err := ws_object_attribute.Delete(deleteDTO.Uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			res.Error = &server.JsonRpcInvalidParamsError
			res.Error.Data = db.ErrNotFound
			return http.StatusNotFound
		}
		logger.Error(err)
		res.Error = &server.JsonRpcInternalError
		return http.StatusInternalServerError
	}

	res.Result = exp
	return http.StatusOK
}

func detailsHandler(req *server.JsonRpcRequest, res *server.JsonRpcResponse) int {
	detailsDTO := &objectGroupAttributeDto.DetailsRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, detailsDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	exp, err := ws_object_attribute.Get(detailsDTO.Uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			res.Error = &server.JsonRpcInvalidParamsError
			res.Error.Data = db.ErrNotFound
			return http.StatusNotFound
		}
		logger.Error(err)
		res.Error = &server.JsonRpcInternalError
		return http.StatusInternalServerError
	}

	res.Result = exp
	return http.StatusOK
}
