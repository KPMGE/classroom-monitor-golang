package httpprotocols

type HttpResponse struct {
	Body       any
	StatusCode int
}

type HttpRequest struct {
	Params any
	Body   any
}

func NewHttpRequest(params any, body any) *HttpRequest {
	return &HttpRequest{
		Params: params,
		Body:   body,
	}
}
