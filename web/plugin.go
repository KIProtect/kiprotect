package web

import (
	"github.com/gospel-sh/gospel"
	"github.com/kiprotect/kodex/api"
)

type WebPluginMaker interface {
	InitializeWebPlugin(controller api.Controller) (WebPlugin, error)
}

type WebPlugin interface {
	MainRoutes(gospel.Context) []*gospel.RouteConfig
}

type AppLink struct {
	Name      string
	Path      string
	Icon      string
	Superuser bool
}

type AppLinkPlugin interface {
	AppLink() AppLink
}

type UserProviderPlugin interface {
	LoginPath() string
}
