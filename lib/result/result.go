package result

type Result struct {
	Status int
	Region string
	Info   string
	Err    error
}
