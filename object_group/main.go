package main

import (
	"github.com/mdalbrid/go-api-aws/object_group/config"
	"github.com/mdalbrid/models/db"
	"github.com/mdalbrid/utils/contexts"
	"github.com/mdalbrid/utils/server"
)

var ctx contexts.InternalModuleMainContext

func init() {
	ctx = contexts.NewInternalModuleMainContext("AWS_ObjectGroupServer")
	db.ReposInit(config.DbConfig)
}

func main() {
	ctx.Logger().Info(" Start")
	jsonRpcHandler := server.NewJsonRpcHandler("/api")

	jsonRpcHandler.MethodHandlerFunc("object_group.create", createHandler)
	jsonRpcHandler.MethodHandlerFunc("object_group.list", listHandler)
	jsonRpcHandler.MethodHandlerFunc("object_group.details", detailsHandler)
	jsonRpcHandler.MethodHandlerFunc("object_group.edit", editHandler)
	jsonRpcHandler.MethodHandlerFunc("object_group.delete", deleteHandler)

	httpServer := server.New(ctx, config.ServerConfig, jsonRpcHandler)
	httpServer.Listen()
}
