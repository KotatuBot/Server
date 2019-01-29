package main

import (
	"./analysis"
	"fmt"
	"time"
)

func main() {

	// geting_list
	//ip_data, mac_data := analysis.GetRegisterApi()

	var data_file string
	data_file = "./Data/http.pcapng"
	var ip_data [2]string
	var mac_data [2]string

	ip_data[0] = "172.21.1.63"
	ip_data[1] = "172.217.161.238"

	mac_data[0] = "00:50:56:c0:00:08"
	mac_data[1] = "6a:00:01:e1:f1:11"

	// infunity loop
	for {
		flag := analysis.Exists(data_file)
		// in Data pcap
		if flag == true {
			//control_ip := "192.168.241.143"
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
			}
		}

	}
	time.Sleep(100 * time.Second)
	fmt.Println("sleep")
}
