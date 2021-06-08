package handlers

import (
	"github.com/jason-shen/clubhouse-clone-biz/config"
	"github.com/jason-shen/clubhouse-clone-biz/ent"
)

type Handler struct {
	Client *ent.Client
	Config *config.Config
}

func NewHandlers(client *ent.Client, config *config.Config) *Handler {
	return &Handler{
		Client: client,
		Config: config,
	}
}
