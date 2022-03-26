package controllers

import (
    "NagoBackend/models"
)

type APIIndexController struct {
    APIBaseController
}

func (base *APIIndexController) Index() {
    var user models.User
    base.DB.First(&user)
    base.OutputJSON(user.Name)
}
