package utils

import "errors"

// Erros específicos da aplicação
var (
	ErrEmailAlreadyExists     = errors.New("EMAIL_ALREADY_EXISTS")
	ErrUserNotFound           = errors.New("USER_NOT_FOUND")
	ErrInvalidCredentials     = errors.New("INVALID_CREDENTIALS")
	ErrRoleNotFound           = errors.New("ROLE_NOT_FOUND")
	ErrRefreshTokenExpired    = errors.New("REFRESH_TOKEN_EXPIRED")
	ErrInvalidPhotoFormat     = errors.New("INVALID_PHOTO_FORMAT")
	ErrInvalidTimestampFormat = errors.New("INVALID_TIMESTAMP_FORMAT")
)

// Função para verificar se é um erro específico
func IsAppError(err error, appErr error) bool {
	return err != nil && err.Error() == appErr.Error()
}

// Função para traduzir erros da aplicação para mensagens do usuário
func TranslateAppError(err error) string {
	switch {
	case IsAppError(err, ErrEmailAlreadyExists):
		return "este e-mail já está em uso"
	case IsAppError(err, ErrUserNotFound):
		return "usuário não encontrado"
	case IsAppError(err, ErrInvalidCredentials):
		return "credenciais inválidas"
	case IsAppError(err, ErrRoleNotFound):
		return "role não encontrado"
	case IsAppError(err, ErrRefreshTokenExpired):
		return "refresh token expirado ou inválido"
	case IsAppError(err, ErrInvalidPhotoFormat):
		return "formato de foto inválido"
	case IsAppError(err, ErrInvalidTimestampFormat):
		return "formato de timestamp inválido"
	default:
		return "erro interno do servidor"
	}
}
