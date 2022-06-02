package presentationprotocols

import httpprotocols "github.com/monitoring-go/src/presentation/http-protocols"

type Controller interface {
	Handle(request *httpprotocols.HttpRequest) *httpprotocols.HttpResponse
}
