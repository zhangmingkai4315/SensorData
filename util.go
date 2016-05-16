package SensorData

import (
	"bufio"
	"os"
	"strings"
)

type SensorDataNetwork struct {
	Sensor_Type   string                    `json:"sensor_type"`
	Sensor_Data   map[string]*InterfaceJson `json:"sensor_data"`
	Error_Message string                    `json:"error_message"`
}

type SensorDataCPU struct {
	Sensor_Type   string         `json:"sensor_type"`
	Sensor_Data   *CPUSensorDate `json:"sensor_data"`
	Error_Message string         `json:"error_message"`
}
type SensorDataMem struct {
	Sensor_Type   string         `json:"sensor_type"`
	Sensor_Data   *MemSensorDate `json:"sensor_data"`
	Error_Message string         `json:"error_message"`
}

func _ReadLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{""}, err
	}
	// 关闭文件
	defer f.Close()

	var ret []string

	r := bufio.NewScanner(f)
	for r.Scan() {
		ret = append(ret, strings.Trim(r.Text(), "\n"))
	}
	return ret, nil
}
