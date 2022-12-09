package sports_test

import (
	"reflect"
	"testing"

	"github.com/plaub/go-test-project/sports"
)

func TestReadSportsData(t *testing.T) {
	sportsData, sportsError := sports.ReadSportsData("../testdata/json.json")

	if sportsError != nil {
		t.Errorf(sportsError.Error())
	}

	expectedLength := 2972
	floatArrsToTest := []string{
		"SpeedInMetersPerSecond",
		"TotalDistanceInMeters",
	}

	for _, property := range floatArrsToTest {
		r := reflect.ValueOf(sportsData)
		f := reflect.Indirect(r).FieldByName(property)
		arr := f.Interface().([]float64)

		if len(arr) != expectedLength {
			t.Errorf("Expected %d, got %d", expectedLength, len(arr))
		}
	}

	intArrsToTest := []string{
		"StartTimeInSeconds",
		"Cadence",
		"HeartRate",
	}

	for _, property := range intArrsToTest {
		r := reflect.ValueOf(sportsData)
		f := reflect.Indirect(r).FieldByName(property)
		arr := f.Interface().([]int64)

		if len(arr) != expectedLength {
			t.Errorf("Expected %d, got %d", expectedLength, len(arr))
		}
	}
}

func TestGetEntriesLength(t *testing.T) {
	sportsData, sportsError := sports.ReadSportsData("../testdata/json.json")

	if sportsError != nil {
		t.Errorf(sportsError.Error())
	}

	expected := 2972
	actual := sportsData.GetEntriesLength()

	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestAverageHeartRate(t *testing.T) {
	sportsData, sportsError := sports.ReadSportsData("../testdata/json.json")

	if sportsError != nil {
		t.Errorf(sportsError.Error())
	}

	expected := int64(103)
	actual := sportsData.GetAverageHeartRate()

	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestAverageSpeedInMetersPerSeconds(t *testing.T) {
	sportsData, sportsError := sports.ReadSportsData("../testdata/json.json")

	if sportsError != nil {
		t.Errorf(sportsError.Error())
	}

	expected := float64(1.168130216159899204)
	actual := sportsData.GetAverageSpeedInMetersPerSeconds()

	if expected != actual {
		t.Errorf("Expected %.18f, got %.18f", expected, actual)
	}
}

func TestGetTotalDurationInSeconds(t *testing.T) {
	sportsData, sportsError := sports.ReadSportsData("../testdata/json.json")

	if sportsError != nil {
		t.Errorf(sportsError.Error())
	}

	expected := int64(2971)
	actual := sportsData.GetTotalDurationInSeconds()

	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetTotalDistanceInMeters(t *testing.T) {
	sportsData, sportsError := sports.ReadSportsData("../testdata/json.json")

	if sportsError != nil {
		t.Errorf(sportsError.Error())
	}

	expected := float64(3722.090000000000145519)
	actual := sportsData.GetTotalDistanceInMeters()

	if expected != actual {
		t.Errorf("Expected %.18f, got %.18f", expected, actual)
	}
}
