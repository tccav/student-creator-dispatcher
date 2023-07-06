package httpclient

import (
	"net/http"
	"net/url"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type DefaultClient struct {
	*http.Client

	baseURL *url.URL
}

func NewDefaultClient(baseURL string) (DefaultClient, error) {
	client := otelhttp.DefaultClient
	client.Timeout = 30 * time.Second

	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return DefaultClient{}, err
	}

	return DefaultClient{
		Client:  client,
		baseURL: parsedURL,
	}, nil
}
