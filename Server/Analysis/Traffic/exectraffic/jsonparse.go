package exectraffic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type TrafficData struct {
	Date          string `json:Date`
	Recive_byte   string `json:Recive_byte`
	Send_byte     string `json:Send_byte`
	Recive_packet string `json:Recive_packet`
	Send_packet   string `json:Send_packet`
}

func JsonParse() {
	bytes, err := ioutil.ReadFile("./Data/write.json")
	if err != nil {

		log.Fatal(err)
	}

	td := new(TrafficData)
	if err := json.Unmarshal(bytes, td); err != nil {
		log.Fatal(err)
	}

	fmt.Println(td.Date)
	fmt.Println(td.Recive_byte)
	fmt.Println(td.Send_byte)
	fmt.Println(td.Recive_packet)
	fmt.Println(td.Send_packet)
	Store(td.Date, td.Recive_byte, td.Send_byte, td.Recive_packet, td.Send_packet)
}
