package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HostLocation struct {
	City        string  `json:"city"`
	RegionCode  string  `json:"region_code"`
	AreaCode    string  `json:"area_code"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
}

type Host struct {
	Asn       string   `json:"asn"`
	Hash      int      `json:"hash"`
	Port      int      `json:"port"`
	Domains   []string `json:"domains"`
	Os        string   `json:"os"`
	Timestamp string   `json:"timestamp"`
	Isp       string   `json:"isp"`
	Transport string   `json:"transport"`
	IpStr     string   `json:"ip_str"`
}
type HostSearch struct {
	Matches []Host `json:"matches"`
}

func (s *Client) HostSearch(q string) (*HostSearch, error) {
	res, err := http.Get(
		fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s", BaseURL, s.apiKey, q),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var ret HostSearch
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
