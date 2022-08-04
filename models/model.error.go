package models

type DatabaseError struct {
	Type string
	Code int
}

type ErrorResponse struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Error      interface{} `json:"error"`
}

type UnathorizatedError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Method  string `json:"method"`
	Message string `json:"message"`
}