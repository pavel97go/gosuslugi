package apperr

import "errors"

// Базовые ошибки приложения, которые мы будем маппить на HTTP-коды.
var (
	ErrValidation = errors.New("validation error") // 400
	ErrNotFound   = errors.New("not found")        // 404
	ErrConflict   = errors.New("conflict")         // 409
)
