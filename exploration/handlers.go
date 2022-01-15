package main

import (
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/mdalbrid/go-api-aws/dto"
	explorationDto "github.com/mdalbrid/go-api-aws/dto/modules/exploration"
	"github.com/mdalbrid/models/db"
	"github.com/mdalbrid/models/ws_exploration"
	"github.com/mdalbrid/utils/logger"
	"github.com/mdalbrid/utils/server"
	"net/http"
)

func listHandler(req *server.JsonRpcRequest, res *server.JsonRpcResponse) int {
	listDTO := &explorationDto.ListRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, listDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	list, err := ws_exploration.List(listDTO.ToFilter())
	if err != nil {
		logger.Error(err)
		res.Error = &server.JsonRpcInternalError
		return http.StatusInternalServerError
	}

	total, err := ws_exploration.Count(listDTO.ToFilter())
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
	createDTO := &explorationDto.CreateRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, createDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	exp, err := ws_exploration.Create(createDTO.Name, createDTO.Comment, createDTO.AccessType, createDTO.Tags)
	if err != nil {
		logger.Error(err)
		res.Error = &server.JsonRpcInternalError
		return http.StatusInternalServerError
	}

	res.Result = exp
	return http.StatusOK
}

func editHandler(req *server.JsonRpcRequest, res *server.JsonRpcResponse) int {
	editDTO := &explorationDto.EditRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, editDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	exp, err := ws_exploration.Edit(editDTO.Uuid, editDTO.Name, editDTO.Comment, editDTO.AccessType, editDTO.Tags)
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
	deleteDTO := &explorationDto.DeleteRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, deleteDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	exp, err := ws_exploration.Delete(deleteDTO.Uuid)
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
	detailsDTO := &explorationDto.DetailsRequestDto{}
	jsonErr := dto.ParseRequestParams(req.Params, detailsDTO)
	if jsonErr != nil {
		res.Error = jsonErr
		return http.StatusBadRequest
	}

	exp, err := ws_exploration.Get(detailsDTO.Uuid)
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
