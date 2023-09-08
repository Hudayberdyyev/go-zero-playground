package domain

type Authorization interface {
	ParseToken(token string) (userId uint32, err error)
}
