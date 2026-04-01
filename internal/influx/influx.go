package influx

import (
	"context"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

type Client struct {
	c      influxdb2.Client
	org    string
	bucket string
}

// Returns the pointer to a Client struct with the influxdb2 client
func New(url, token, org, bucket string) *Client {
	return &Client{
		c:      influxdb2.NewClient(url, token),
		org:    org,
		bucket: bucket,
	}
}

// Apply the flux query and returns an interface
func (i *Client) Query(ctx context.Context, flux string) ([]map[string]interface{}, error) {
	queryAPI := i.c.QueryAPI(i.org)
	result, err := queryAPI.Query(ctx, flux)
	if err != nil {
		return nil, err
	}

	rows := make([]map[string]interface{}, 0)
	for result.Next() {
		rec := result.Record()
		row := rec.Values()

		rows = append(rows, row)
	}
	if result.Err() != nil {
		return nil, result.Err()
	}

	return rows, nil
}
