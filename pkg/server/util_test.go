package server

import "testing"

func TestgetTemperatureData(t *testing.T) {
	for i, test := range []struct {
		jsonStr []byte
		want    TemperatureData
	}{
		{
			jsonStr: []byte("{\"DataSource\": \"Something\", \"OutdoorTemp\": 4.0, \"IndoorTemp\": 20.0}"),
			want: TemperatureData{
				OutdoorTemp: 4.0,
				IndoorTemp:  20.0,
			},
		},
		{
			jsonStr: []byte("{\"DataSource\": \"Something\", \"OutdoorTemp\": 4.0}"),
			want: TemperatureData{
				OutdoorTemp: 4.0,
			},
		},
		{
			jsonStr: []byte("{\"RandomField\": \"Something\", \"OutdoorTemp\": 4.0}"),
			want: TemperatureData{
				OutdoorTemp: 4.0,
			},
		},
	} {
		data := getTemperatureData(test.jsonStr)

		if !data.DataAreEqual(test.want) {
			t.Errorf("Test #%d: Got\n%v\nWanted\n%v\n", i, data, test.want)
		}
	}
}
