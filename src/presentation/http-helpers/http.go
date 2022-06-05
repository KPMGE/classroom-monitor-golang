package httphelpers

import httpprotocols "github.com/monitoring-go/src/presentation/http-protocols"

func ServerError(err error) *httpprotocols.HttpResponse {
	return &httpprotocols.HttpResponse{
		Body:       err,
		StatusCode: 500,
	}
}

func Ok(body any) *httpprotocols.HttpResponse {
	return &httpprotocols.HttpResponse{
		Body:       body,
		StatusCode: 200,
	}
}

func BadRequest(message string) *httpprotocols.HttpResponse {
	return &httpprotocols.HttpResponse{
		Body:       message,
		StatusCode: 400,
	}
}
