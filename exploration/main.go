package main

import (
	"github.com/mdalbrid/go-api-aws/exploration/config"
	"github.com/mdalbrid/models/db"
	"github.com/mdalbrid/utils/contexts"
	"github.com/mdalbrid/utils/server"
)

var ctx contexts.InternalModuleMainContext

func init() {
	ctx = contexts.NewInternalModuleMainContext("AWS_ExplorationServer")
	db.ReposInit(config.DbConfig)
}

func main() {
	ctx.Logger().Info(" Start")
	jsonRpcHandler := server.NewJsonRpcHandler("/api")

	jsonRpcHandler.MethodHandlerFunc("exploration.list", listHandler)
	jsonRpcHandler.MethodHandlerFunc("exploration.create", createHandler)
	jsonRpcHandler.MethodHandlerFunc("exploration.edit", editHandler)
	jsonRpcHandler.MethodHandlerFunc("exploration.details", detailsHandler)
	jsonRpcHandler.MethodHandlerFunc("exploration.delete", deleteHandler)

	httpServer := server.New(ctx, config.ServerConfig, jsonRpcHandler)
	httpServer.Listen()
}
