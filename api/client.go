package api

import (
	"io"
	"net/http"
	"time"

	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
)

var client *httpclient.Client

func init() {
	initialTimeout := 2 * time.Millisecond
	maxTimeout := 9 * time.Millisecond
	exponentFactor := 2
	maximumJitterInterval := 2 * time.Millisecond

	backoff := heimdall.NewExponentialBackoff(initialTimeout, maxTimeout, float64(exponentFactor), maximumJitterInterval)
	retrier := heimdall.NewRetrier(backoff)
	timeout := 1000 * time.Millisecond

	client = httpclient.NewClient(
		httpclient.WithHTTPTimeout(timeout),
		httpclient.WithRetrier(retrier),
		httpclient.WithRetryCount(4),
	)
}

func getRequest(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	return body, nil
}
