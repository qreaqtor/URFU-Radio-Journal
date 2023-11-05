package controllers

import (
	"urfu-radio-journal/pkg/services"
)

type EditionController struct {
	es services.EditionService
}

func NewEditionController() EditionController {
	return EditionController{es: services.NewEditionService()}
}
