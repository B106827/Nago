package controllers

import (
    "net/http"
)

type response struct {
    Status  int         `json:"status"`
    Result  interface{} `json:"result"`
}

func successResponse(msg interface{}) *response {
    return &response{http.StatusOK, msg}
}

func badRequestResponse(msg []string) *response {
    return &response{http.StatusBadRequest, msg}
}
