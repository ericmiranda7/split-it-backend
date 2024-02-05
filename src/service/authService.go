package service

type AuthService struct {
}

func GetAuthService() *AuthService {
	return &AuthService{}
}

func (as *AuthService) getCreds() {
}
