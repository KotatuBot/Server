package analysis

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func Split() {
	fmt.Println("start")
	cmd := exec.Command("editcap", "-c", "100", "/opt/Data/http.pcapng", "/opt/Split/test.pcap")
	cmd.Start()
	cmd.Wait()
	fmt.Println("end")

}

// get dir data
func GetList() []string {
	var file_list []string
	list, err := ioutil.ReadDir("/opt/Split")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	for _, file := range list {
		path := "/opt/Split/" + file.Name()
		file_list = append(file_list, path)
	}

	return file_list
}
