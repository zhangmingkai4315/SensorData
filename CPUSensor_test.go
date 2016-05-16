package SensorData

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

var fileName string

func BeforeTest() string {
	content := []byte("0.20 0.30 0.20 32/269 11067")
	tmpfile, err := ioutil.TempFile("", "cpu_test_file")
	if err != nil {
		panic(err)
	}
	if _, err := tmpfile.Write(content); err != nil {
		panic(err)
	}
	if err := tmpfile.Close(); err != nil {
		panic(err)
	}

	return tmpfile.Name()
}

func AfterTest() {
	os.Remove(fileName)
}
func TestGetCPUData(t *testing.T) {
	fileName = BeforeTest()
	output, err := GetCpuData(fileName)
	if err != nil {
		t.Errorf("Expected data message, but %v was instead", err)
	}
	data := new(SensorDataCPU)
	t.Log("Test if function will return the rigth CPU data")
	if err != nil {
		t.Errorf("Expected data message, but error was instead")
	} else {
		err := json.Unmarshal(output, &data)
		if err != nil {
			panic(err)
		}
		if data.Sensor_Data.OneMin != 0.2 {
			t.Errorf("Expected OneMin data == 0.2, but %v was instead", data.Sensor_Data.OneMin)
		}
		t.Logf("Test Data is %+v", data)
	}
	AfterTest()
}
