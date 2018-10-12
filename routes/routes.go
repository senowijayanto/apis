package routes

import (
	"github.com/gorilla/mux"
	"github.com/senowijayanto/apis/controllers"
)

type Route struct{}

func (r *Route) Init() *mux.Router {
	macroController := controllers.InitMacroController()

	router := mux.NewRouter().StrictSlash(false)
	v1 := router.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/macros", macroController.GetListMacro).Methods("GET")

	return v1
}
