package server

import (
	"encoding/json"
	"math"
	"time"
)

type TemperatureData struct {
	Timestamp   string
	DataSource  string
	OutdoorTemp float64
	IndoorTemp  float64
}

// DataAreEqual check if all data fields (e.g. not Timestamp) in two structures match
func (td TemperatureData) DataAreEqual(other TemperatureData) bool {
	tol := 1e-6
	return (td.DataSource == other.DataSource) && (math.Abs(td.OutdoorTemp-other.OutdoorTemp) < tol) &&
		(math.Abs(td.IndoorTemp-other.IndoorTemp) < tol)
}

func GetTemperatureData(jsonData []byte) TemperatureData {
	data := TemperatureData{}
	json.Unmarshal(jsonData, &data)

	data.Timestamp = time.Now().String()
	return data
}
