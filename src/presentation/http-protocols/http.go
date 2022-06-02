package httpprotocols

type HttpResponse struct {
	Body       any
	StatusCode int
}

type HttpRequest struct {
	Params any
	Body   any
}
