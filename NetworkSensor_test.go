package SensorData

import "testing"
import "encoding/json"

func TestGetNetworkData(t *testing.T) {
	output := new(SensorDataOutput)
	t.Log("Test if function will return the rigth data")
	b, err := GetNetworkInfo(2, "*", "")
	if err != nil {
		t.Errorf("Expected data message, but error was instead")
		return
	}
	if err := json.Unmarshal(b, &output); err != nil {
		panic(err)
	}
	if output.Sensor_Type != "NetworkTrafficSensor" || len(output.Sensor_Data) <= 0 {
		t.Errorf("Expected data message, but 0 interface was returned")
	}
}

func TestNonExistFile(t *testing.T) {
	t.Log("Test if function will return the rigth error message when file is non-exist")
	_, err := GetNetworkInfo(2, "*", "non-exist.file")
	if err == nil {
		t.Errorf("Expected error message, but nil was instead")
		return
	}
}
