package getpcap

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	//interface_name string = "enx3495db11dfe2"
	interface_name string = "wlp3s0"

	Write_File string = "/opt/Data/http.pcapng"
	Time_out   string = "duration:10"
)

func GetPcap() {
	cmd := exec.Command("tshark", "-i", interface_name, "-a", Time_out, "-w", Write_File)

	cmd.Start()
	cmd.Wait()

	if err := os.Chmod(Write_File, 0777); err != nil {
		fmt.Println(err)
	}
}
