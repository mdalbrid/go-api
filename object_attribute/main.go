package main

import (
	"github.com/mdalbrid/go-api-aws/object_attribute/config"
	"github.com/mdalbrid/models/db"
	"github.com/mdalbrid/utils/contexts"
	"github.com/mdalbrid/utils/server"
)

var ctx contexts.InternalModuleMainContext

func init() {
	ctx = contexts.NewInternalModuleMainContext("AWS_ObjectAttributeServer")
	db.ReposInit(config.DbConfig)
}

func main() {
	ctx.Logger().Info(" Start")
	jsonRpcHandler := server.NewJsonRpcHandler("/api")

	jsonRpcHandler.MethodHandlerFunc("object_attribute.create", createHandler)
	jsonRpcHandler.MethodHandlerFunc("object_attribute.list", listHandler)
	jsonRpcHandler.MethodHandlerFunc("object_attribute.details", detailsHandler)
	jsonRpcHandler.MethodHandlerFunc("object_attribute.edit", editHandler)
	jsonRpcHandler.MethodHandlerFunc("object_attribute.delete", deleteHandler)

	httpServer := server.New(ctx, config.ServerConfig, jsonRpcHandler)
	httpServer.Listen()
}
