package analysis

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

func test() {
	fmt.Println(url.Parse("http://localhost:8000/api/traffic"))
}
func GenerateControl(url, Control_ip, Connect_ip, Country, City, Sessions, Date string) error {
	Date2 := "54"
	fmt.Println("control")
	jsonStr := `{"Control_ip":"` + Control_ip + `","Connect_ip":"` + Connect_ip + `","Country":"` + Country + `","City":"` + City + `","Sessions":"` + Sessions + `","Date":"` + Date2 + `"}`
	fmt.Println(jsonStr)

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	fmt.Println(req)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println(resp)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err

}

func GeneratePort(url, Connect_ip, Port, Port_Sessions string) error {

	jsonStr := `{"Connect_ip":"` + Connect_ip + `","Port":"` + Port + `","Port_Sessions":"` + Port_Sessions + `"}`

	//fmt.Println(url.Parse("http://localhost:8000/api/traffic"))

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println(resp)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err

}

func GenerateTraffic(url, ip, macaddress, date, recv_byte, recv_packet, send_byte, send_packet string) error {

	//fmt.Println(url.Parse("http://localhost:8000/api/traffic"))

	jsonStr := `{"ip":"` + ip + `","macaddress":"` + macaddress + `","date":"` + date + `","recv_byte":"` + recv_byte + `","recv_packet":"` + recv_packet + `","send_byte":"` + send_byte + `","send_packet":"` + send_packet + `"}`

	fmt.Println(jsonStr)
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println(resp)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err

}
