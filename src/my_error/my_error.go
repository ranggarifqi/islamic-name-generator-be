package my_error

type Error interface {
	Error() string
	GetStatusCode() int
}
