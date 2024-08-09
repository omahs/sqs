package passthroughdomain

import (
	"net/http"

	"github.com/osmosis-labs/sqs/sqsutil/sqshttp"
)

type TimeSeriesHTTPClient interface {
	GetPoolFees() (*PoolFees, error)
}

type TimeSeriesHTTPClientImpl struct {
	url    string
	client *http.Client
}

const (
	feesEndpoint = "/fees/v1/pools"
)

var _ TimeSeriesHTTPClient = &TimeSeriesHTTPClientImpl{}

func NewTimeSeriesHTTPClient(url string) *TimeSeriesHTTPClientImpl {
	return &TimeSeriesHTTPClientImpl{
		url:    url,
		client: &http.Client{},
	}
}

// GetPoolFees implements TimeSeriesHTTPClient.
func (t *TimeSeriesHTTPClientImpl) GetPoolFees() (*PoolFees, error) {
	poolFees, err := sqshttp.Get[PoolFees](t.client, t.url, feesEndpoint)
	if err != nil {
		return nil, err
	}
	return poolFees, nil
}
