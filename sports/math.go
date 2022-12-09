package sports

// CalculateAverageSpeed calculates the average speed in meters per second.
func CalculateAverageSpeed(distanceInMeters float64, durationInSeconds float64) (meterPerSecond float64) {
	return distanceInMeters / durationInSeconds
}

// CalculatePace calculates the pace in minutes per kilometer.
func CalculatePace(speedInMeterPerSecond float64) (minPerKilometer float64) {
	if speedInMeterPerSecond == 0 {
		return 0
	}

	return 1 / (speedInMeterPerSecond / 1000) / 60
}

// GetPaceFormatted returns the pace in minutes per kilometer as a formatted string.
func GetPaceFormatted(speedInMeterPerSecond float64) (paceFormatted string) {
	return SecondsToTimeString(int64(CalculatePace(speedInMeterPerSecond) * 60))
}
