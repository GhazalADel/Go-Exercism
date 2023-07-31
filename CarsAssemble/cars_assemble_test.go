package cars

import (
	"math"
	"testing"
)

const FloatEqualityThreshold = 1e-9

func TestCalculateProductionRatePerHour(t *testing.T) {
	tests := []struct {
		name           string
		productionRate int
		successRate    float64
		want           float64
	}{
		{
			name:           "calculate working cars per hour for production rate 0",
			productionRate: 0,
			successRate:    100,
			want:           0.0,
		},
		{
			name:           "calculate working cars per hour for 100% success rate",
			productionRate: 221,
			successRate:    100,
			want:           221.0,
		},
		{
			name:           "calculate working cars per hour for 80% success rate",
			productionRate: 426,
			successRate:    80,
			want:           340.8,
		},
		{
			name:           "calculate working cars per hour for 20.5% success rate",
			productionRate: 6824,
			successRate:    20.5,
			want:           1398.92,
		}, {
			name:           "calculate working cars per hour for 0% success rate",
			productionRate: 8000,
			successRate:    0,
			want:           0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateWorkingCarsPerHour(tt.productionRate, tt.successRate)
			if math.Abs(got-tt.want) > FloatEqualityThreshold {
				t.Errorf(
					"CalculateProductionRatePerHour(%d,%f) = %f, want %f",
					tt.productionRate,
					tt.successRate,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestCalculateProductionRatePerMinute(t *testing.T) {
	tests := []struct {
		name           string
		productionRate int
		successRate    float64
		want           int
	}{
		{
			name:           "calculate working cars per minute for production rate 0",
			productionRate: 0,
			successRate:    100,
			want:           0,
		},
		{
			name:           "calculate working cars per minute for 100% success rate",
			productionRate: 221,
			successRate:    100,
			want:           3,
		},
		{
			name:           "calculate working cars per minute for 80% success rate",
			productionRate: 426,
			successRate:    80,
			want:           5,
		},
		{
			name:           "calculate working cars per minute for 20.5% success rate",
			productionRate: 6824,
			successRate:    20.5,
			want:           23,
		},
		{
			name:           "calculate working cars per minute for 0% success rate",
			productionRate: 8000,
			successRate:    0,
			want:           0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateWorkingCarsPerMinute(tt.productionRate, tt.successRate)
			if got != tt.want {
				t.Errorf(
					"CalculateWorkingItemsRate(%d,%f) = %d, want %d",
					tt.productionRate,
					tt.successRate,
					got,
					tt.want,
				)
			}
		})
	}
}
