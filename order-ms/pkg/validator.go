package pkg

import (
	"order-ms/internal/domain"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateRequest(req any) *domain.AppError {
    err := validate.Struct(req)
    if err == nil {
        return nil
    }

    // If it's a validation error, process it
    if ve, ok := err.(validator.ValidationErrors); ok {
        var messages []string
        for _, fieldErr := range ve {
            msg := formatValidationError(fieldErr)
            messages = append(messages, msg)
        }
        return domain.NewAppError(strings.Join(messages, ", "), "VALIDATION_ERROR", http.StatusBadRequest)
    }

    return domain.BadRequestErr("invalid request")
}

func formatValidationError(fe validator.FieldError) string {
    field := fe.Field()
    tag := fe.Tag()

    switch tag {
    case "required":
        return field + " is required"
    case "email":
        return field + " must be a valid email"
    case "min":
        return field + " must be at least " + fe.Param() + " characters"
    case "max":
        return field + " must be at most " + fe.Param() + " characters"
    case "uuid4", "uuid":
        return field + " must be a valid UUID"
    default:
        return field + " is not valid"
    }
}
