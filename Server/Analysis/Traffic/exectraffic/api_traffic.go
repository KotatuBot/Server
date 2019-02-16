package exectraffic

import (
	"bytes"
	"fmt"
	"net/http"
)

func HttpPost(url, date, recv_byte, recv_packet, send_byte, send_packet string) error {

	jsonStr := `{"date":"` + date + `","recv_byte":"` + recv_byte + `","recv_packet":"` + recv_packet + `","send_byte":"` + send_byte + `","send_packet":"` + send_packet + `"}`

	fmt.Println(jsonStr)

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		fmt.Println("fff")
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("sss")
		return err
	}
	defer resp.Body.Close()

	return err

}
