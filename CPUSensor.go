package SensorData

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type CPUSensorDate struct {
	OneMin         float64 `json:"one_min_load"`
	FiveMin        float64 `json:"five_min_load"`
	TenMin         float64 `json:"ten_min_load"`
	RunningProcess float64 `json:"running_processes"`
	TotalProcess   float64 `json:"total_processes"`
}

func SensorMessageFmt(data *SensorDataCPU, err error) ([]byte, error) {
	data.Error_Message = fmt.Sprintf("%v", err)
	output, _ := json.Marshal(data)
	return output, err
}

func parseFloat64(s string, b int) float64 {
	output, err := strconv.ParseFloat(s, b)
	if err != nil {
		return -1
	} else {
		return output
	}
}

func GetCpuData(devfile string) ([]byte, error) {

	sensor_output := &SensorDataCPU{
		Sensor_Type:   "CPUSensor",
		Sensor_Data:   &CPUSensorDate{},
		Error_Message: "",
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	lines, err := _ReadLines(devfile)
	if err != nil {
		return SensorMessageFmt(sensor_output, err)
	}
	fields := strings.Split(lines[0], " ")
	processes := strings.Split(fields[3], "/")
	c := CPUSensorDate{
		OneMin:         parseFloat64(fields[0], 64),
		FiveMin:        parseFloat64(fields[1], 64),
		TenMin:         parseFloat64(fields[2], 64),
		RunningProcess: parseFloat64(processes[0], 64),
		TotalProcess:   parseFloat64(processes[1], 64),
	}
	sensor_output.Sensor_Data = &c
	return SensorMessageFmt(sensor_output, nil)
}
