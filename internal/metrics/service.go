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
	return s.repo.QueryEnergyData(ctx, "today()", "now()", "")
}

// Returs a slice of EnergergyPoints of yesterday
func (s *Service) GetYesterdayEnergy(ctx context.Context) ([]EnergyPoint, error) {
	return s.repo.QueryEnergyData(ctx, "yesterday()", "today()", "")
}

// Returs a slice of EnergergyPoints of a set range
func (s *Service) GetRangeEnergyByDay(ctx context.Context, start, stop time.Time) ([]DailyEnergy, error) {
	points, err := s.repo.QueryEnergyData(ctx, start.Format(time.RFC3339), stop.Format(time.RFC3339), "10m")
	if err != nil {
		return nil, err
	}
	return groupByDay(points), nil
}

// Devides the slice of EnergyPoints per each day in the struct DailyEnergy
func groupByDay(points []EnergyPoint) []DailyEnergy {
	byDay := make(map[string][]EnergyPoint)

	for _, p := range points {
		day := p.Time.Format("206-01-02")
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
