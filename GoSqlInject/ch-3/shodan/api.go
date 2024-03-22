package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIInfo struct {
	ScanCredits  int `json:"scan_credits"`
	UsageLimits  `json:"usage_limits"`
	Plan         string      `json:"plan"`
	Https        bool        `json:"https"`
	Unlocked     bool        `json:"unlocked"`
	QueryCredits int         `json:"query_credits"`
	MonitoredIps interface{} `json:"monitored_ips"`
	UnlockedLeft int         `json:"unlocked_left"`
	Telnet       bool        `json:"telnet"`
}

type UsageLimits struct {
	ScanCredits  int `json:"scan_credits"`
	QueryCredits int `json:"query_credits"`
	MonitoredIps int `json:"monitored_ips"`
}

func (s *Client) APIinfo() (*APIInfo, error) {
	res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, s.apiKey))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var ret APIInfo
	//返回一个带焕缓存的解码器，然后可以进行输入，然后调用方法对输入进行解码
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
