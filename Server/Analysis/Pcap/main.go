package main

import (
	"./analysis"
	"fmt"
	"os"
	"time"
)

func main() {

	// geting_list
	ip_data, mac_data := analysis.GetRegisterApi()
	fmt.Println(ip_data)

	var data_file string
	data_file = "/opt/Data/http.pcapng"

	// infunity loop
	for {
		flag := analysis.Exists(data_file)
		// in Data pcap
		if flag == true {
			fmt.Println("True")
			//split file
			analysis.Split()

			//get file_list
			file_list := analysis.GetList()
			for _, file_path := range file_list {
				// analysis
				for number, control_ip := range ip_data {
					analysis.Pcap_Master(control_ip, mac_data[number], file_path)
					fmt.Println("sucess")

				}
				if err := os.Remove(file_path); err != nil {
					fmt.Println("delete")
				}
			}
			if err := os.Remove(data_file); err != nil {
				fmt.Println("delete")
			}
		}

		time.Sleep(5 * time.Second)
	}
}
