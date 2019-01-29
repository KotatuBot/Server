package analysis

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"time"
)

func Parse_connect(control_ip string, session_data map[string]int) {

	db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/Packets")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for key, _ := range session_data {
		fmt.Println("Control IP")
		fmt.Println(control_ip)
		fmt.Println("Connect IP")
		fmt.Println(key)
		fmt.Println("Sessions")
		fmt.Println(session_data[key])
		city, country := Geo_search(key)
		fmt.Println(city)
		fmt.Println(country)
		timer := time.Now().String()
		fmt.Println(timer)

		result, err := db.Exec("INSERT INTO IoT_connect(Control_ip,Connect_ip,Country,City,Sessions,Date) VALUES(?,?,?,?,?,?)", control_ip, key, country, city, session_data[key], timer)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(result)
	}
}

func Parse_port(ip_port_data map[string]int) {

	db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/Packets")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for key, _ := range ip_port_data {
		split_data := strings.Split(key, ",")
		fmt.Println(split_data[0])
		fmt.Println(split_data[1])
		session_number := strconv.Itoa(ip_port_data[key])
		fmt.Println(session_number)

		result, err := db.Exec("INSERT INTO IoT_port(Connect_ip,Port,Port_Sessions) VALUES(?,?,?)", split_data[0], split_data[1], session_number)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(result)
	}

}

func Parse_traffic(control_ip string, mac_address string, send_byte int, recv_byte int, send_packet int, recv_packet int) {
	db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/Packets")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Traffic")
	fmt.Println(control_ip)
	fmt.Println(mac_address)
	fmt.Println(send_byte)
	fmt.Println(recv_byte)
	timer := time.Now().String()
	fmt.Println(timer)

	result, err := db.Exec("INSERT INTO IoT_eachtraffic(ip,macaddress,date,recv_byte,send_byte,recv_packet,send_packet) VALUES(?,?,?,?,?,?,?)", control_ip, mac_address, timer, recv_byte, send_byte, recv_packet, send_packet)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(result)
}
