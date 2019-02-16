package analysis

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Parse_connect(control_ip string, session_data map[string]int) {

	for key, _ := range session_data {
		fmt.Println("Control IP")
		fmt.Println(control_ip)
		fmt.Println("Connect IP")
		fmt.Println(key)
		fmt.Println("Sessions")
		fmt.Println(session_data[key])
		sessions := strconv.Itoa(session_data[key])
		city, country := Geo_search(key)
		fmt.Println(city)
		fmt.Println(country)
		timer := time.Now().Format("2016-01-03")
		fmt.Println(timer)

		GenerateControl("http://localhost:8000/api/connect/", control_ip, key, country, city, sessions, timer)
		fmt.Println("---------")

	}
}

func Parse_port(ip_port_data map[string]int) {

	for key, _ := range ip_port_data {
		split_data := strings.Split(key, ",")
		fmt.Println(split_data[0])
		fmt.Println(split_data[1])
		session_number := strconv.Itoa(ip_port_data[key])
		fmt.Println(session_number)

		GeneratePort("http://localhost:8000/api/ports/", split_data[0], split_data[1], session_number)

	}

}

func Parse_traffic(control_ip string, mac_address string, send_byte int, recv_byte int, send_packet int, recv_packet int) {

	fmt.Println("Traffic")
	fmt.Println(control_ip)
	fmt.Println(mac_address)
	fmt.Println(send_byte)
	fmt.Println(recv_byte)
	timers := time.Now().UTC()
	timer := timers.Format("2006-01-02T15:04:05Z")

	send_bytes := strconv.Itoa(send_byte)
	recv_bytes := strconv.Itoa(recv_byte)
	send_packets := strconv.Itoa(send_packet)
	recv_packets := strconv.Itoa(recv_packet)

	GenerateTraffic("http://localhost:8000/api/eachtraffic/", control_ip, mac_address, timer, recv_bytes, send_bytes, recv_packets, send_packets)
}
