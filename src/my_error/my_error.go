package my_error

type MyError interface {
	Error() string
	GetStatusCode() int
}
