package exectraffic

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Store(Date string, Recive_byte string, Send_byte string, Recive_packet string, Send_packet string) {
	db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/Packets")
	if err != nil {

		panic(err.Error())
	}

	defer db.Close()

	result, err := db.Exec("INSERT INTO IoT_traffic(date,recv_byte,recv_packet,send_byte,send_packet) VALUES(?,?,?,?,?)", Date, Recive_byte, Recive_packet, Send_byte, Send_packet)

	if err != nil {

		panic(err.Error())
	}
	fmt.Println(result)

}
