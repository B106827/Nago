package controllers

import (
    "net/http"
)

type response struct {
    Status  int         `json:"status"`
    Result  interface{} `json:"result"`
}

// 200
func successResponse(res interface{}) *response {
    return &response{http.StatusOK, res}
}

// 400
func badRequestResponse(res interface{}) *response {
    return &response{http.StatusBadRequest, res}
}

// 401
func unauthorizedResponse(res interface{}) *response {
    return &response{http.StatusUnauthorized, res}
}

// 500
func serverErrorResponse(res interface{}) *response {
    return &response{http.StatusInternalServerError, res}
}
