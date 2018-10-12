package services

import (
	"github.com/senowijayanto/apis/infrastructures"
	"github.com/senowijayanto/apis/models"
	"github.com/senowijayanto/apis/repositories"
)

type IMacroService interface {
	GetListMacro() (models.MacroResponse, error)
}

type MacroService struct {
	MacroRepository repositories.IMacroRepository
}

func InitMacroService() *MacroService {
	macroRepository := new(repositories.MacroRepository)
	macroRepository.DB = &infrastructures.MYSQLConnection{}

	macroService := new(MacroService)
	macroService.MacroRepository = macroRepository

	return macroService
}

func (m *MacroService) GetListMacro() (macros models.MacroResponse, err error) {
	macros, err = m.MacroRepository.GetListMacro()
	return
}
