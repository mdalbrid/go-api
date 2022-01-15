package main

import (
	"github.com/mdalbrid/go-api-aws/object/config"
	"github.com/mdalbrid/models/db"
	"github.com/mdalbrid/utils/contexts"
	"github.com/mdalbrid/utils/server"
)

var ctx contexts.InternalModuleMainContext

func init() {
	ctx = contexts.NewInternalModuleMainContext("AWS_ObjectServer")
	db.ReposInit(config.DbConfig)
}

func main() {
	ctx.Logger().Info(" Start")
	jsonRpcHandler := server.NewJsonRpcHandler("/api")

	jsonRpcHandler.MethodHandlerFunc("object.create", createHandler)
	jsonRpcHandler.MethodHandlerFunc("object.list", listHandler)
	jsonRpcHandler.MethodHandlerFunc("object.details", detailsHandler)
	jsonRpcHandler.MethodHandlerFunc("object.edit", editHandler)
	jsonRpcHandler.MethodHandlerFunc("object.delete", deleteHandler)

	httpServer := server.New(ctx, config.ServerConfig, jsonRpcHandler)
	httpServer.Listen()
}
