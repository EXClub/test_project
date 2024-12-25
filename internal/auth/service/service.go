package service

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) GenerateHash() {
}

func (s *AuthService) CompareHash() {
}

func (s *AuthService) GenerateToken() {}

func (s *AuthService) CompareToken() {

}
