package controllers

import (
	"net/http"

	"github.com/senowijayanto/apis/helpers"
	"github.com/senowijayanto/apis/services"
)

func InitMacroController() *MacroController {
	macroService := services.InitMacroService()

	macroController := new(MacroController)
	macroController.MacroService = macroService
	return macroController
}

type MacroController struct {
	MacroService services.IMacroService
}

func (c *MacroController) GetListMacro(res http.ResponseWriter, req *http.Request) {
	macros, _ := c.MacroService.GetListMacro()
	helpers.Response(res, http.StatusOK, macros)
}
