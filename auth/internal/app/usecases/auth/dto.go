package auth

type (
	RegistrationDTO struct {
		Username string
		Password string
	}

	LoginDTO struct {
		Username string
		Password string
	}

	LogoutDTO struct {
		Username string
	}

	VerifyDTO struct {
		Username string
		Token    string
	}
)
