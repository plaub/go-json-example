package sports_test

import (
	"testing"

	"github.com/plaub/go-json-example/sports"
)

func TestCalculatePace(t *testing.T) {
	tests := []struct {
		avgSpeedInMeterPerSecond float64
		expectedPace             float64
	}{
		{5.8312603648, 2.858158549611992338},
		{4.5, 3.703703703703703720},
	}

	for _, test := range tests {
		pace := sports.CalculatePace(test.avgSpeedInMeterPerSecond)
		if pace != test.expectedPace {
			t.Errorf("Expected %.18f, got %.18f", test.expectedPace, pace)
		}
	}
}

func TestCalculateAverageSpeed(t *testing.T) {
	tests := []struct {
		distanceInMeters float64
		durationInHours  float64
		expectedSpeed    float64
	}{
		{42195, 2.01, 5.831260364842455246},
		{1000, 1, 0.277777777777777790},
		{10000, 1, 2.777777777777777679},
	}

	for _, test := range tests {
		speed := sports.CalculateAverageSpeed(test.distanceInMeters, test.durationInHours*3600)
		if speed != test.expectedSpeed {
			t.Errorf("Expected %.18f, got %.18f", test.expectedSpeed, speed)
		}
	}
}

func TestTestCalculatePaceFromAverageSpeed(t *testing.T) {
	tests := []struct {
		distanceInMeters float64
		durationInHours  float64
		expectedPace     float64
	}{
		{42195, 2.0, 2.843938855314610770},
	}

	for _, test := range tests {
		pace := sports.CalculatePace(sports.CalculateAverageSpeed(test.distanceInMeters, test.durationInHours*3600))
		if pace != test.expectedPace {
			t.Errorf("Expected %.18f, got %.18f", test.expectedPace, pace)
		}
	}
}

func TestGetPaceFormatted(t *testing.T) {
	tests := []struct {
		avgSpeedInMeterPerSecond float64
		expectedPace             string
	}{
		{5.8312603648, "00:02:51"},
		{2.1340089631388928, "00:07:48"},
	}

	for _, test := range tests {
		pace := sports.GetPaceFormatted(test.avgSpeedInMeterPerSecond)
		if pace != test.expectedPace {
			t.Errorf("Expected %s, got %s", test.expectedPace, pace)
		}
	}
}
