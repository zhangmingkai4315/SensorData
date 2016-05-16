package SensorData

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type MemSensorDate struct {
	MemTotal   float64 `json:"memory_total"`
	MemFree    float64 `json:"memory_free"`
	MemUsed    float64 `json:"memory_used"`
	Buffers    float64 `json:"buffers"`
	Cached     float64 `json:"cached"`
	SwapCached float64 `json:"swap_cached"`
	SwapTotal  float64 `json:"swap_total"`
	SwapFree   float64 `json:"swap_free"`
}

func (data *SensorDataMem) SensorMessageFmt(err error) ([]byte, error) {
	data.Error_Message = fmt.Sprintf("%v", err)
	output, _ := json.Marshal(data)
	return output, err
}
func _GetKeyValue(s string) float64 {
	splited := strings.Split(s, " ")
	return parseFloat64(splited[0], 64)
}
func GetMemData(devfile string) ([]byte, error) {
	if devfile == "" {
		devfile = "/proc/meminfo"
	}
	sensor_output := &SensorDataMem{
		Sensor_Type:   "MemSensor",
		Sensor_Data:   &MemSensorDate{},
		Error_Message: "",
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	lines, err := _ReadLines(devfile)
	if err != nil {
		return sensor_output.SensorMessageFmt(err)
	}
	result_map := make(map[string]string)
	for _, line := range lines {
		field := strings.Split(line, ":")
		result_map[field[0]] = strings.TrimSpace(field[1])
	}
	c := MemSensorDate{
		MemTotal:   _GetKeyValue(result_map["MemTotal"]),
		MemFree:    _GetKeyValue(result_map["MemFree"]),
		MemUsed:    _GetKeyValue(result_map["MemTotal"]) - _GetKeyValue(result_map["MemFree"]),
		Buffers:    _GetKeyValue(result_map["Buffers"]),
		Cached:     _GetKeyValue(result_map["Cached"]),
		SwapCached: _GetKeyValue(result_map["SwapCached"]),
		SwapTotal:  _GetKeyValue(result_map["SwapTotal"]),
		SwapFree:   _GetKeyValue(result_map["SwapFree"]),
	}
	sensor_output.Sensor_Data = &c
	return sensor_output.SensorMessageFmt(nil)
}
