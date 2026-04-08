package metrics

import (
	"context"
	"fmt"
	"time"

	"github.com/chiaf1/solar-api/internal/influx"
)

// Specific queries for data retrive from db

type Repository struct {
	influx *influx.Client
}

func NewRepository(client *influx.Client) *Repository {
	return &Repository{influx: client}
}

func (r *Repository) QueryDailyData(ctx context.Context, fromt, to time.Time) ([]map[string]interface{}, error) {
	flux := fmt.Sprintf(`from(bucket: "%s")
	|> range(start: today(), stop: now())
	|> filter(fn: (r) => r._measurement == "energy")
	|> filter(fn: (r) => r._field == "production" or r._field == "consumption")
	|> sort(columns: ["_time"])`, r.influx.Bucket)
	return r.influx.Query(context.Background(), flux)
}
