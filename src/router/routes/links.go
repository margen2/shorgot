package routes

import (
	"net/http"

	"github.com/margen2/shorgot/src/controllers"
)

var linkRoutes = []Route{
	{
		URI:                    "/links",
		Method:                 http.MethodPost,
		Function:               controllers.CreateLink,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/links/{link}",
		Method:                 http.MethodGet,
		Function:               controllers.SearchLink,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/links",
		Method:                 http.MethodGet,
		Function:               controllers.SearchLinks,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/links/{linkID}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateLink,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/links/{linkID}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteLink,
		RequiresAuthentication: true,
	},
}
