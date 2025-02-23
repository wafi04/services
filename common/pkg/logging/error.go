package logging

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Custom error constants
const (
	ErrorUnauthorized   = "Unauthorized access"
	ErrorNotFound       = "Resource not found"
	ErrorBadRequest     = "Bad request"
	ErrorInternalServer = "Internal server error"
)

// Error codes
const (
	CodeUnauthorized   = 401
	CodeNotFound       = 404
	CodeBadRequest     = 400
	CodeInternalServer = 500
)

// CreateError membuat error response
func CreateError(code int, message string) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: message,
	}
}

// Contoh penggunaan untuk error spesifik
var (
	ErrUnauthorized   = CreateError(CodeUnauthorized, ErrorUnauthorized)
	ErrNotFound       = CreateError(CodeNotFound, ErrorNotFound)
	ErrBadRequest     = CreateError(CodeBadRequest, ErrorBadRequest)
	ErrInternalServer = CreateError(CodeInternalServer, ErrorInternalServer)
)
