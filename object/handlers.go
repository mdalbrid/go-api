package main

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v4"

	"github.com/mdalbrid/go-api-aws/dto"
	objectDto "github.com/mdalbrid/go-api-aws/dto/modules/object"
	"github.com/mdalbrid/models/db"
	"github.com/mdalbrid/models/ws_object"
	"github.com/mdalbrid/utils/logger"
	"github.com/mdalbrid/utils/server"
)

func listHandler(req *server.JsonRpcRequest, res *server.JsonRpcResponse) int {
	listDTO := &objectDto.ListRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, listDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	list, err := ws_object.List(listDTO.ToFilter())
	if err != nil {
		logger.Error(err)
		res.Error = &server.JsonRpcInternalError
		return http.StatusInternalServerError
	}

	total, err := ws_object.Count(listDTO.ToFilter())
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
	createDTO := &objectDto.CreateRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, createDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	exp, err := ws_object.Create(createDTO.ExplorationUUID, createDTO.GroupUUID, createDTO.Name, createDTO.Image, createDTO.Country, createDTO.Comment)
	if err != nil {
		logger.Error(err)
		res.Error = &server.JsonRpcInternalError
		return http.StatusInternalServerError
	}

	res.Result = exp
	return http.StatusOK
}

func editHandler(req *server.JsonRpcRequest, res *server.JsonRpcResponse) int {
	editDTO := &objectDto.EditRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, editDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	exp, err := ws_object.Edit(editDTO.Uuid, editDTO.GroupUUID, editDTO.Name, editDTO.Image, editDTO.Country, editDTO.Comment, editDTO.SortWeight)
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
	deleteDTO := &objectDto.DeleteRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, deleteDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	exp, err := ws_object.Delete(deleteDTO.Uuid)
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
	detailsDTO := &objectDto.DetailsRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, detailsDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	exp, err := ws_object.Get(detailsDTO.Uuid)
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
