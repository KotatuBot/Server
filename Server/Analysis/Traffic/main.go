package main

import (
	"./exectraffic"
	"fmt"
	"os"
	"time"
)

func main() {
	var file_name string
	file_name = "./Data/write.json"
	// ifunity loop
	for {
		result := exectraffic.Exists(file_name)

		if result == false {
			fmt.Println("No")
			time.Sleep(5 * time.Second)
		} else {
			exectraffic.JsonParse()
			if err := os.Remove(file_name); err != nil {
				fmt.Println(err)
			}

		}
	}
}
