package conf

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
)

type HandlerSettings struct {
	NodeDirPath              string `json:"node_path"`
	ClientExecutableFullPath string `json:"executable_path"`
	HTTPInterfaceHost        string `json:"http_host"`
	HTTPInterfacePort        uint16 `json:"http_port"`
}

type ServiceSettings struct {
	EventsMonitorExecutableFullPath string `json:"events_monitor_path"`
	SendEvents                      bool   `json:"allow_send_events"`
	SendLogs                        bool   `json:"allow_send_logs"`
	Host                            string `json:"host"`
	Port                            uint16 `json:"port"`
}

type SecuritySettings struct {
	ApiKey       string   `json:"api_key"`
	AllowableIPs []string `json:"allowable_ips"`
}

type Settings struct {
	Handler  HandlerSettings  `json:"handler"`
	Service  ServiceSettings  `json:"collecting_data_service"`
	Security SecuritySettings `json:"security"`
}

func (s ServiceSettings) ServiceInterface() string {
	return "http://" + s.Host + ":" + strconv.Itoa(int(s.Port))
}

func (s HandlerSettings) HTTPInterface() string {
	return s.HTTPInterfaceHost + ":" + strconv.Itoa(int(s.HTTPInterfacePort))
}

var (
	Params = Settings{}
)

func LoadSettings() error {
	// Reading configuration file
	bytes, err := ioutil.ReadFile("conf.json")
	if err != nil {
		return errors.New("Can't read configuration. Details: " + err.Error())
	}

	err = json.Unmarshal(bytes, &Params)
	if err != nil {
		return errors.New("Can't read configuration. Details: " + err.Error())
	}

	return nil
}
