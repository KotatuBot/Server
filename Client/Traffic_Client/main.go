package main

import (
	"./gettraffic"
	"fmt"
	"os/exec"
	"time"
)

func send(ipaddress string, data string) {
	_, err := exec.Command("python", "client.py", ipaddress, data).Output()
	if err != nil {
		panic(err)
	}

}

func main() {
	fmt.Println("Traffic")
	for value := 0; ; {
		value--
		time.Sleep(17 * time.Second)
		gettraffic.GetTraffic()
		send("127.0.0.1", "traffic")
	}
}
