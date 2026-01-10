package index

import (
	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/pages/index/handlers"
)

func GetHandlers() []*common.Handler {
	return []*common.Handler{
		common.NewHandler("/", "GET", handlers.IndexHandler),
		common.NewHandler("/home", "GET", handlers.IndexHandler),
	}
}
