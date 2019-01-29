package analysis

import (
	"github.com/google/gopacket"
	"strconv"
	"strings"
)

func Packet_Byte(packet gopacket.Packet) int {
	var byte_size int
	frame_data := packet.String()
	later_data := strings.Split(frame_data, "PACKET: ")[1]
	bytes := strings.Split(later_data, " bytes")[0]
	byte_size, _ = strconv.Atoi(bytes)
	return byte_size

}

func Get_Byte(control_ip string, src string, dst string, packet gopacket.Packet) (int, int, int, int) {
	var byte_packet_size int
	var send_byte int
	var recv_byte int
	var send_packet int
	var recv_packet int

	if control_ip == src {
		byte_packet_size = Packet_Byte(packet)
		send_byte = byte_packet_size
		recv_byte = 0

		send_packet = 1
		recv_packet = 0

	} else if control_ip == dst {

		byte_packet_size = Packet_Byte(packet)
		recv_byte = byte_packet_size
		send_byte = 0

		recv_packet = 1
		send_packet = 0

	} else {
		send_byte = 0
		recv_byte = 0

		send_packet = 0
		recv_packet = 1
	}

	return send_byte, recv_byte, send_packet, recv_packet
}
