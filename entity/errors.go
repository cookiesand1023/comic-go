package errors

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

var DuplicateEmailError = Error{
	Message: "email already exists",
	Code:    1001,
}

var EmailNotFound = Error{
	Message: "email not found",
	Code:    1002,
}

var PasswordInvalid = Error{
	Message: "password is invalid",
	Code:    1003,
}
