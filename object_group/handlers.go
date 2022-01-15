package main

import (
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/mdalbrid/go-api-aws/dto"
	objectGroupDto "github.com/mdalbrid/go-api-aws/dto/modules/object_group"
	"github.com/mdalbrid/models/db"
	"github.com/mdalbrid/models/ws_object_group"
	"github.com/mdalbrid/utils/logger"
	"github.com/mdalbrid/utils/server"
	"net/http"
)

func listHandler(req *server.JsonRpcRequest, res *server.JsonRpcResponse) int {
	listDTO := &objectGroupDto.ListRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, listDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	list, err := ws_object_group.List(listDTO.ToFilter())
	if err != nil {
		logger.Error(err)
		res.Error = &server.JsonRpcInternalError
		return http.StatusInternalServerError
	}

	total, err := ws_object_group.Count(listDTO.ToFilter())
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
	createDTO := &objectGroupDto.CreateRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, createDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	exp, err := ws_object_group.Create(createDTO.ExplorationUUID, createDTO.Name, createDTO.Icon)
	if err != nil {
		logger.Error(err)
		res.Error = &server.JsonRpcInternalError
		return http.StatusInternalServerError
	}

	res.Result = exp
	return http.StatusOK
}

func editHandler(req *server.JsonRpcRequest, res *server.JsonRpcResponse) int {
	editDTO := &objectGroupDto.EditRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, editDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	exp, err := ws_object_group.Edit(editDTO.Uuid, editDTO.Name, editDTO.Icon, editDTO.SortWeight)
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
	deleteDTO := &objectGroupDto.DeleteRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, deleteDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	exp, err := ws_object_group.Delete(deleteDTO.Uuid)
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
	detailsDTO := &objectGroupDto.DetailsRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, detailsDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	exp, err := ws_object_group.Get(detailsDTO.Uuid)
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
