package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/eddyhub/wifictl/system"
)

var defaultMux *mux.Router

func init() {
	defaultMux = mux.NewRouter()
}

func setDefaultServerMux(mux *mux.Router)  {
	defaultMux = mux
}

func SetRoutes(mux *mux.Router) {
	if mux != nil {
		setDefaultServerMux(mux)
	}

	mux.HandleFunc("/api/isEnabled", startServiceHandle).Methods(http.MethodPut)
	mux.HandleFunc("/api/isEnabled", stopServiceHandle).Methods(http.MethodDelete)
}

func stopServiceHandle(w http.ResponseWriter, r *http.Request) {
	system.StopHostapd()
}

func startServiceHandle(w http.ResponseWriter, r *http.Request) {
	system.StartHostapd()
}
