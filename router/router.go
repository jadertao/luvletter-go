package router

import (
	"luvletter/app/mood"
	"luvletter/app/tag"

	"github.com/labstack/echo"

	"luvletter/app/letter"
	"luvletter/app/user"
)

// Prefix ...
var Prefix = "/api/v1"

// PrefixMapper ...
func PrefixMapper(router map[string]echo.HandlerFunc, prefix string) map[string]echo.HandlerFunc {
	var withPrefixRouter = make(map[string]echo.HandlerFunc)
	for key, value := range router {
		withPrefixRouter[prefix+key] = value
	}
	return withPrefixRouter
}

// GETRouters RouterConfig for GET.
var GETRouters = PrefixMapper(map[string]echo.HandlerFunc{
	"/letter": letter.GetAll,
}, Prefix)

// POSTRouters RouterConfig for POST.
var POSTRouters = PrefixMapper(map[string]echo.HandlerFunc{
	"/login":    user.Login,
	"/register": user.Register,
	"/letter":   letter.Save,
	"/tag":      tag.Save,
	"/mood":     mood.Save,
}, Prefix)
