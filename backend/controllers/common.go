package controllers

import (
	"net/http"
)

type response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

type cveResponse struct {
	Status   int         `json:"status"`
	Result   interface{} `json:"result"`
	IsCveErr bool        `json:"isCustomValidErr"`
}

// 200
func successResponse(res interface{}) *response {
	return &response{http.StatusOK, res}
}

// 400（カスタムバリデーションエラー）
func customValidErrResponse(res interface{}) *cveResponse {
	return &cveResponse{http.StatusBadRequest, res, true}
}

// 400
func badRequestResponse(res interface{}) *response {
	return &response{http.StatusBadRequest, res}
}

// 401
func unauthorizedResponse(res interface{}) *response {
	return &response{http.StatusUnauthorized, res}
}

// 404
func notFoundResponse(res interface{}) *response {
	return &response{http.StatusNotFound, res}
}

// 500
func serverErrorResponse(res interface{}) *response {
	return &response{http.StatusInternalServerError, res}
}
