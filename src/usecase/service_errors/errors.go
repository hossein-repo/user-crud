// usecase/service_errors/errors.go
package service_errors

type ServiceError struct {
    EndUserMessage string
    TechnicalError error
    ErrorCode      int
}

func (e *ServiceError) Error() string {
    return e.EndUserMessage
}

var (
    EmailExists        = "email already exists"
    UsernameExists     = "username already exists"
    MobileExists       = "mobile number already exists"
    UserNotFound       = "user not found"
    InvalidCredentials = "invalid credentials"
    TokenRequired      = "token required"
    TokenExpired       = "token expired"
    TokenInvalid       = "token invalid"
)
