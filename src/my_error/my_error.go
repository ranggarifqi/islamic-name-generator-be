package my_error

type Error interface {
	error
	GetStatusCode() int
}
