package metrics

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/chiaf1/solar-api/internal/influx"
)

// Specific queries for data retrive from db

type Repository struct {
	influx *influx.Client
}

type EnergyPoint struct {
	Time        time.Time `json:"time"`
	Production  float64   `json:"production"`
	Consumption float64   `json:"consumption"`
}

// Creates a repository for data queries
func NewRepository(client *influx.Client) *Repository {
	return &Repository{influx: client}
}

// Query production and consumption from the db
// start and stop is the range for the query
// window is the agregation window, if left empty it wont be used (format 5m, )
func (r *Repository) QueryEnergyData(ctx context.Context, start, stop, window string) ([]EnergyPoint, error) {
	flux := fmt.Sprintf(`from(bucket: "%s")
	|> range(start: %s, stop: %s)
	|> filter(fn: (r) => r._measurement == "energy")
	|> filter(fn: (r) => r._field == "production" or r._field == "consumption")`, r.influx.Bucket, start, stop)

	if window != "" {
		flux += fmt.Sprintf(`
		|> aggregateWindow(every:%s, fn: mean, createEmpty: false)
		`, window)
	}

	flux += `|> sort(columns: ["_time"])`

	results, err := r.influx.Query(context.Background(), flux)
	if err != nil {
		return nil, err
	}

	// The issue here is that the return of the query flux is vertical so for the same time stamp
	// the value productoin and consumptio are separated in two different points so we need to reconnect them

	// Create a map of energy point serialized on time.Time to create the []EnergyPoint struct
	pointMap := make(map[time.Time]*EnergyPoint)

	// Parse the results to store in the pointMap
	for _, row := range results {

		// Parsing the time from string to time.Time
		/* tsStr, ok := row["_time"].(string)
		if !ok {
			continue
		}
		t, err := time.Parse(time.RFC3339, tsStr)
		if err != nil {
			continue
		} */
		t := row["_time"].(time.Time)

		// Create or reuse the EnergyPoint in the map
		point, exists := pointMap[t]
		if !exists {
			point = &EnergyPoint{Time: t}
			pointMap[t] = point
		}

		// Now we parse the _field name and value and store the correct one
		field, _ := row["_field"].(string)
		val, _ := row["_value"].(float64)

		switch field {
		case "production":
			point.Production = val
		case "consumption":
			point.Consumption = val
		}
	}

	// Convert the map in slice to return it
	points := make([]EnergyPoint, 0, len(pointMap))
	for _, p := range pointMap {
		points = append(points, *p)
	}

	// Sort the slice on time
	sort.Slice(points, func(i, j int) bool {
		return points[i].Time.Before(points[j].Time)
	})

	return points, nil
}
