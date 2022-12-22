package worker

// RequestDto
// Modify your api_request data here
type RequestDto struct {
	Description string
}

type Worker struct {
	I           int
	Description string
	Duration    int64
	Err         error
}
