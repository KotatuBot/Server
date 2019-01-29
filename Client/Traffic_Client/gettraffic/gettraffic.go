package gettraffic

import (
	"os/exec"
)

func GetTraffic() {

	_, err := exec.Command("python", "./gettraffic/traffic.py").Output()
	if err != nil {
		panic(err)
	}

}
