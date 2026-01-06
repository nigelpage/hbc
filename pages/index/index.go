package index

import (
	i "github.com/nigelpage/hbc/internal"
	h "github.com/nigelpage/hbc/pages/index/handlers"
)

func GetHandlers() []i.Handler {
	return []i.Handler{
		{UrlPattern: "/",		 				Verb: "GET", 	Function: h.IndexHandler},
		{UrlPattern: "/home",		 			Verb: "GET", 	Function: h.IndexHandler},
	}
}
