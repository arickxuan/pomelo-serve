package route

import (
	"pomeloServe/connector/handler"
	"pomeloServe/core/repo"
	"pomeloServe/framework/net"
)

func Register(r *repo.Manager) net.LogicHandler {
	handlers := make(net.LogicHandler)
	entryHandler := handler.NewEntryHandler(r)
	handlers["entryHandler.entry"] = entryHandler.Entry

	return handlers
}
