package metrics

import (
	"context"
	"sort"
	"time"
)

// Queries' results manipulation and data formatting

type Service struct {
	repo *Repository
}

type DailyEnergy struct {
	Day    string        `json:"day"`
	Points []EnergyPoint `json:"points"`
}

// Returns the pointer to a new Service struct
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// Returs a slice of EnergergyPoints of today
func (s *Service) GetTodayEnergy(ctx context.Context) ([]EnergyPoint, error) {
	now := time.Now().UTC()

	startDay := truncateToDay(now)

	return s.repo.QueryEnergyData(ctx, startDay.Format(time.RFC3339), now.Format(time.RFC3339), "2m")
}

// Returs a slice of EnergergyPoints of yesterday
func (s *Service) GetYesterdayEnergy(ctx context.Context) ([]EnergyPoint, error) {
	now := time.Now().UTC()

	startDay := truncateToDay(now.AddDate(0, 0, -1))
	stopDay := truncateToDay(now)

	return s.repo.QueryEnergyData(ctx, startDay.Format(time.RFC3339), stopDay.Format(time.RFC3339), "2m")
}

// Returs a slice of EnergergyPoints of a set range
func (s *Service) GetRangeEnergyByDay(ctx context.Context, start, stop time.Time, window string) ([]DailyEnergy, error) {
	if window == "" {
		window = "10m"
	}

	startDay := truncateToDay(start)
	stopDay := truncateToDay(stop).Add(24 * time.Hour)

	points, err := s.repo.QueryEnergyData(ctx, startDay.Format(time.RFC3339), stopDay.Format(time.RFC3339), window)
	if err != nil {
		return nil, err
	}
	return groupByDay(points), nil
}

// Devides the slice of EnergyPoints per each day in the struct DailyEnergy
func groupByDay(points []EnergyPoint) []DailyEnergy {
	byDay := make(map[string][]EnergyPoint)

	for _, p := range points {
		day := p.Time.Format("2006-01-02")
		byDay[day] = append(byDay[day], p)
	}

	days := make([]string, 0, len(byDay))
	for day := range byDay {
		days = append(days, day)
	}

	sort.Strings(days)

	out := make([]DailyEnergy, 0, len(byDay))
	for _, day := range days {
		out = append(out, DailyEnergy{
			Day:    day,
			Points: byDay[day],
		})
	}

	return out
}

// Truncate time value to start of day
func truncateToDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}
