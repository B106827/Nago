package controllers

import (
    "net/http"
)

type response struct {
    Status  int         `json:"status"`
    Result  interface{} `json:"result"`
}

func successResponse(res interface{}) *response {
    return &response{http.StatusOK, res}
}

func badRequestResponse(res interface{}) *response {
    return &response{http.StatusBadRequest, res}
}
