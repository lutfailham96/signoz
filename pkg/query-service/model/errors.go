package model

// custom errors related to registration

type ErrEmailRequired struct{}

func (errEmailRequired ErrEmailRequired) Error() string {
	return "email is required"
}

type ErrPasswordRequired struct{}

func (errPasswordRequired ErrPasswordRequired) Error() string {
	return "password is required"
}

type ErrSignupFailed struct{}

func (errSignupFailed ErrSignupFailed) Error() string {
	return "failed to register user"
}

type ErrNoOrgFound struct{}

func (errNoOrgFound ErrNoOrgFound) Error() string {
	return "no org found"
}
