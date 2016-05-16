package SensorData

import (
	"encoding/json"
	"testing"
)

func TestGetMemData(t *testing.T) {

	output, err := GetMemData("./example/meminfo")
	if err != nil {
		t.Errorf("Expected data message, but %v was instead", err)
	}
	data := new(SensorDataMem)
	t.Log("Test if function will return the rigth Memory data")
	if err != nil {
		t.Errorf("Expected Memory message, but error was instead")
	} else {
		err := json.Unmarshal(output, &data)
		if err != nil {
			panic(err)
		}
		if data.Sensor_Data.MemTotal != 1939840 {
			t.Errorf("Expected OneMin data == 1939840, but %v was instead", data.Sensor_Data.MemTotal)
		}
		t.Logf("Test Data is %+v", data)
	}
	AfterTest()
}
