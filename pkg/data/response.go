package data

type Response struct {
	Success bool
	Data    string
}

func FailedResponse(msg string) Response {
	return Response{Success: false, Data: msg}
}
