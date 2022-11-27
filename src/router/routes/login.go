package routes

import (
	"github.com/margen2/shorgot/src/controllers"
	"net/http"
)

var loginRoute = Route{
	URI:                    "/login",
	Method:                 http.MethodPost,
	Function:               controllers.Login,
	RequiresAuthentication: false,
}
