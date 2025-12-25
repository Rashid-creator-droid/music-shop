package auth

type MockAuthService struct {
	RegisterCalled bool
	LoginCalled    bool
	Token          string
	Err            error
}

func (m *MockAuthService) Register(input RegisterRequest) error {
	m.RegisterCalled = true
	return m.Err
}

func (m *MockAuthService) Login(input LoginRequest) (string, error) {
	m.LoginCalled = true
	return m.Token, m.Err
}
