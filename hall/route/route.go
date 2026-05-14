package route

import (
	"pomeloServe/core/repo"
	"pomeloServe/framework/node"
	"pomeloServe/hall/handler"
)

func Register(r *repo.Manager) node.LogicHandler {
	handlers := make(node.LogicHandler)
	userHandler := handler.NewUserHandler(r)
	handlers["userHandler.updateUserAddress"] = userHandler.UpdateUserAddress
	return handlers
}
