package sports

import (
	"fmt"

	"github.com/rodaine/table"
)

// SportsDataTablePrint prints a table with the given sports data.
func SportsDataTablePrint(sportsData []SportsData) {
	tbl := table.New("#", "Duration", "Distance (km)", "Avg HR", "Avg Pace (min/km)", "Avg Speed (m/s)")

	for index, item := range sportsData {
		totalTimeFormatted := SecondsToTimeString(item.GetTotalDurationInSeconds())
		avgSpeedInMetersPerSeconds := item.GetAverageSpeedInMetersPerSeconds()
		avgPaceFormatted := GetPaceFormatted(avgSpeedInMetersPerSeconds)
		avgHeartRate := item.GetAverageHeartRate()
		totalDistanceFormatted := fmt.Sprintf("%.2f", item.GetTotalDistanceInMeters()/1000)

		tbl.AddRow(
			index+1,
			totalTimeFormatted,
			totalDistanceFormatted,
			avgHeartRate,
			avgPaceFormatted,
			avgSpeedInMetersPerSeconds,
		)
	}

	tbl.Print()
}
