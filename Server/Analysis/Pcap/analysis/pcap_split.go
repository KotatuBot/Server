package analysis

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func Split() {
	fmt.Println("start")
	cmd := exec.Command("editcap", "-c", "100", "./Data/http.pcapng", "./Split/test.pcap")
	cmd.Start()
	cmd.Wait()
	fmt.Println("end")

}

// get dir data
func GetList() []string {
	var file_list []string
	list, err := ioutil.ReadDir("./Split")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	for _, file := range list {
		path := "./Split/" + file.Name()
		file_list = append(file_list, path)
	}

	return file_list
}
