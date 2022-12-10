package sports

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/plaub/go-test-project/helper"
)

type SportsData struct {
	StartTimeAsDateTime     []string
	StartTimeInSeconds      []int64
	LatitudeInDegrees       []float64
	LongitudeInDegrees      []float64
	ElevationInMeters       []float64
	AirTemperatureCelcius   []int64
	HeartRate               []int64
	SpeedInMetersPerSecond  []float64
	TotalDistanceInMeters   []float64
	TimerDurationInSeconds  []float64
	ClockDurationInSeconds  []float64
	MovingDurationInSeconds []float64
	PowerInWatts            []int64
	Cadence                 []int64
	Calories                []int64
}

func (sd *SportsData) GetEntriesLength() int {
	return len(sd.StartTimeInSeconds)
}

func (sd *SportsData) GetAverageHeartRate() int64 {
	return helper.GetArrayValuesAverage(sd.HeartRate)
}

func (sd *SportsData) GetAverageSpeedInMetersPerSeconds() float64 {
	return helper.GetArrayValuesAverage(sd.SpeedInMetersPerSecond)
}

func (sd *SportsData) GetTotalDurationInSeconds() int64 {
	if len(sd.StartTimeInSeconds) == 0 {
		return 0
	}

	return sd.StartTimeInSeconds[sd.GetEntriesLength()-1] - sd.StartTimeInSeconds[0]
}

func (sd *SportsData) GetTotalDistanceInMeters() float64 {
	if len(sd.TotalDistanceInMeters) == 0 {
		return 0
	}

	return sd.TotalDistanceInMeters[sd.GetEntriesLength()-1]
}

// ReadSportsData reads a json file and returns a SportsData struct.
func ReadSportsData(path string) (SportsData, error) {
	var sportsData SportsData

	// Open our jsonFile
	jsonFile, openError := os.Open(path)
	// if we os.Open returns an error then handle it
	if openError != nil {
		return sportsData, errors.New(openError.Error())
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return sportsData, errors.New(err.Error())
	}

	json.Unmarshal(byteValue, &sportsData)
	return sportsData, nil
}
