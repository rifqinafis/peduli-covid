package messages

import "errors"

var (
	// error_messages
	ErrIDNotFound               = errors.New("id not found")
	ErrDuplicateData            = errors.New("duplicate data")
	ErrDataAlreadyExist         = errors.New("data already exist")
	ErrInternalServer           = errors.New("something gone wrong, contact administrator")
	ErrUsernamePasswordNotFound = errors.New("(Username) or (Password) empty")
	ErrNotFound                 = errors.New("data not found")
	ErrInvalidBearerToken       = errors.New("invalid bearer token")
	ErrExpiredToken             = errors.New("expired token")
	ErrInvalidRole              = errors.New("invalid role")
	ErrInvalidCred              = errors.New("invalid credential")
	ErrInvalidParam             = errors.New("invalid param")

	//Modular
	BaseResponseMessageSuccess = "success"
	BaseResponseMessageFailed  = "something not right"
)
