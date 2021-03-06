package server

import (
	"bytes"
	"encoding/json"
	"math"
	"time"
)

type TemperatureData struct {
	Timestamp   string
	OutdoorTemp float64
	IndoorTemp  float64
	Forecast    float64
}

// DataAreEqual check if all data fields (e.g. not Timestamp) in two structures match
func (td TemperatureData) DataAreEqual(other TemperatureData) bool {
	tol := 1e-6
	return (math.Abs(td.OutdoorTemp-other.OutdoorTemp) < tol) &&
		(math.Abs(td.IndoorTemp-other.IndoorTemp) < tol)
}

func getTemperatureData(jsonData []byte) TemperatureData {
	data := TemperatureData{}
	json.Unmarshal(jsonData, &data)

	data.Timestamp = time.Now().String()
	return data
}

func jsonDataOk(jsonData []byte) bool {
	expectKeys := []string{"OutdoorTemp", "IndoorTemp", "Forecast"}
	for _, key := range expectKeys {
		if !bytes.Contains(jsonData, []byte(key)) {
			return false
		}
	}
	return true
}
