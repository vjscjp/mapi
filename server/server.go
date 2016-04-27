package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vjscjp/mapi/controllers"
	"github.com/vjscjp/mapi/controllers/app"
	"github.com/vjscjp/mapi/controllers/host"
)

const (
	Default   = "/"
	App       = "/app/{id}"
	HostPorts = "/hostport"
	HostPort  = "/hostport/{id}/{port}"
)

func InitRoutes(muxRouter *mux.Router) {
	controllers.SetEnv()
	muxRouter.HandleFunc(Default, Status).Methods("GET")
	muxRouter.HandleFunc(App, appcontroller.GetApp).Methods("GET")
	muxRouter.HandleFunc(HostPorts, hostcontroller.GetHostPorts).Methods("GET")
	muxRouter.HandleFunc(HostPort, hostcontroller.GetHostPort).Methods("GET")
}

func Status(w http.ResponseWriter, r *http.Request) {
	controllers.ServeJsonResponseWithCode(w, map[string]string{"Status": "Ok"}, http.StatusOK)
}
