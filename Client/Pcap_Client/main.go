package main

import (
	"./getpcap"
	//"./gettraffic"
	"fmt"
	"os/exec"
	//"time"
)

func send(ipaddress string, data string) {
	_, err := exec.Command("python", "client.py", ipaddress, data).Output()
	if err != nil {
		panic(err)
	}

}

func main() {
	fmt.Println("Pcap")
	for value := 0; ; {
		value--
		//go func() {
		getpcap.GetPcap()
		send("127.0.0.1", "pcap")

		//}()
		//getpcap.GetPcap()
		//gettraffic.GetTraffic()
		//send("127.0.0.1", "pcap")
		//send("127.0.0.1", "traffic")
	}
}
