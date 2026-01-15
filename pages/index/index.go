package index

import (
	"net/http"

	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/pages/index/handlers"
)

func GetHeaderMenusAndHandlers() *[]common.HeaderMenuAndHandler {
	hmahs :=[]common.HeaderMenuAndHandler{
		{Url: "/", Method: http.MethodGet, Handler: handlers.IndexHandler},
		{Url: "/home", Text: "home", Method: http.MethodGet, Handler: handlers.IndexHandler},
	}
	return &hmahs
}
