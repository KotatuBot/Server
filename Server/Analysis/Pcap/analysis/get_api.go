package analysis

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Control struct {
	Control_ip string `json:"Control_ip"`
	MacAddress string `json:"MacAddress"`
	OS         string `json:"OS"`
	Port       string `json:"Port"`
	Service    string `json:"Service"`
}

func GetRegisterApi() ([]string, []string) {

	var control []Control
	var ip_list []string
	var mac_list []string

	url := "http://localhost:8000/api/control/"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(byteArray, &control); err != nil {
		log.Fatal(err)
	}

	for _, c := range control {

		ip_list = append(ip_list, c.Control_ip)
		mac_list = append(mac_list, c.MacAddress)

	}

	return ip_list, mac_list
}
