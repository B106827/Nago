package controllers

import (
    "NagoBackend/server/contexts"
    "net/http"
)

type APIBaseController struct {
    contexts.CustomContext
}

func (base *APIBaseController) OutputJSON(data interface{}) error {
    return base.JSON(http.StatusOK, data)
}
