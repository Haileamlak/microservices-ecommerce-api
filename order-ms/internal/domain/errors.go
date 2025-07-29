package domain

import "net/http"

type AppError struct {
    Message    string `json:"message"`
    StatusCode int    `json:"-"`
    Code       string `json:"code"` // e.g. "PRODUCT_NOT_FOUND", etc.
}

func (e *AppError) Error() string {
    return e.Message
}

func NewAppError(message, code string, statusCode int) *AppError {
    return &AppError{
        Message:    message,
        Code:       code,
        StatusCode: statusCode,
    }
}

// Common helpers
func NotFoundErr(message string) *AppError {
    return NewAppError(message, "NOT_FOUND", http.StatusNotFound)
}

func BadRequestErr(message string) *AppError {
    return NewAppError(message, "BAD_REQUEST", http.StatusBadRequest)
}

func InternalErr(message string) *AppError {
    return NewAppError("internal server error", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
}

func UnauthorizedErr(message string) *AppError {
	return NewAppError(message, "UNAUTHORIZED", http.StatusUnauthorized)
}

func ForbiddenErr(message string) *AppError {
	return NewAppError(message, "FORBIDDEN", http.StatusForbidden)
}

func ConflictErr(message string) *AppError {
	return NewAppError(message, "CONFLICT", http.StatusConflict)
}

func BadGatewayErr(message string) *AppError {
	return NewAppError(message, "BAD_GATEWAY", http.StatusBadGateway)
}

func ServiceUnavailableErr(message string) *AppError {
	return NewAppError(message, "SERVICE_UNAVAILABLE", http.StatusServiceUnavailable)
}
