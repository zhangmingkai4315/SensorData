package SensorData

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// 定义网络接口数据
type NetStat struct {
	Dev  []string
	Stat map[string]*DevStat
}

// 定义接口的名称和数据量
type DevStat struct {
	Name string
	Rx   uint64
	Tx   uint64
}

type InterfaceJson struct {
	Name string `json:"name"`
	Rx   string `json:"RxSpeed"`
	Tx   string `json:"TxSpeed"`
}

func _GetInfo(filename string, interface_ string) (ret NetStat, err error) {
	lines, err := _ReadLines(filename)
	if err != nil {
		return ret, err
	}

	ret.Dev = make([]string, 0)
	ret.Stat = make(map[string]*DevStat)

	for _, line := range lines {
		// 跳过前面的两行
		fields := strings.Split(line, ":")
		if len(fields) < 2 {
			continue
		}
		key := strings.TrimSpace(fields[0])
		value := strings.Fields(strings.TrimSpace(fields[1]))

		// 根据参数显示全部或者只显示某个接口
		if interface_ != "*" && interface_ != key {
			continue
		}
		c := new(DevStat)
		//		c := DevStat{}
		c.Name = key
		r, err := strconv.ParseInt(value[0], 10, 64)
		if err != nil {
			log.Print(key, "Rx", value[0], err)
			break
		}
		c.Rx = uint64(r)

		t, err := strconv.ParseInt(value[8], 10, 64)
		if err != nil {
			log.Println(key, "Tx", value[8], err)
			break
		}
		c.Tx = uint64(t)
		ret.Dev = append(ret.Dev, key)
		ret.Stat[key] = c
	}
	return
}
func _Vsize(bytes uint64, delta float64) (ret string) {
	var tmp float64 = float64(bytes) / delta
	var s string = " "
	bytes = uint64(tmp)
	switch {
	case bytes < uint64(2<<9):

	case bytes < uint64(2<<19):
		tmp = tmp / float64(2<<9)
		s = "K"

	case bytes < uint64(2<<29):
		tmp = tmp / float64(2<<19)
		s = "M"

	case bytes < uint64(2<<39):
		tmp = tmp / float64(2<<29)
		s = "G"

	case bytes < uint64(2<<49):
		tmp = tmp / float64(2<<39)
		s = "T"

	}
	ret = fmt.Sprintf("%.2f%sB/s", tmp, s)
	return
}
func GetNetworkInfo(interval_time float64, interface_ string, LinuxNetworkFile string) ([]byte, error) {
	const (
		count_num = 1
	)
	if interval_time == 0 {
		interval_time = 3
	}
	if interface_ == "" {
		interface_ = "*"
	}
	if LinuxNetworkFile == "" {
		LinuxNetworkFile = "/tmp/dev"
	}
	log.SetFlags(log.Ldate | log.Ltime)
	sensor_output := SensorDataNetwork{
		Sensor_Type:   "NetworkTrafficSensor",
		Sensor_Data:   make(map[string]*InterfaceJson),
		Error_Message: "",
	}

	var stat0 NetStat
	var delta NetStat
	delta.Dev = make([]string, 0)
	delta.Stat = make(map[string]*DevStat)

	i := count_num
	if i > 0 {
		i += 1
	}

	for {
		stat1, err := _GetInfo(LinuxNetworkFile, interface_)
		if err != nil {
			return []byte{}, err
		}
		for _, value := range stat1.Dev {
			t0, ok := stat0.Stat[value]
			if ok {
				dev, ok := delta.Stat[value]
				if !ok {
					delta.Stat[value] = new(DevStat)
					dev = delta.Stat[value]
					delta.Dev = append(delta.Dev, value)
				}
				t1 := stat1.Stat[value]
				dev.Rx = t1.Rx - t0.Rx
				dev.Tx = t1.Tx - t0.Tx
			}
		}
		stat0 = stat1
		for _, iface := range delta.Dev {
			stat := delta.Stat[iface]
			rx, tx := _Vsize(stat.Rx, interval_time), _Vsize(stat.Tx, interval_time)
			tempJson := InterfaceJson{Name: iface, Rx: rx, Tx: tx}
			sensor_output.Sensor_Data[iface] = &tempJson
		}

		i -= 1
		if i == 0 {
			break
		} else {
			// 休息*秒后执行
			time.Sleep(time.Duration(interval_time*1000) * time.Millisecond)
		}
	}
	output, err := json.Marshal(sensor_output)
	if err != nil {
		return []byte{}, err
	}
	return output, nil
}
