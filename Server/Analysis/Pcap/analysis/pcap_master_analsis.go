package analysis

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"io"
	"log"
)

var (
	handle *pcap.Handle
	err    error
)

func Pcap_Master(control_ip string, mac_address string, pcapFile string) {

	var sets string
	var dst_ip []string
	var dst_ip_port []string
	var unique_ip_dst []string
	var unique_ip_port_dst []string
	var ip_count map[string]int
	var ip_port_count map[string]int
	var send_byte int
	var recv_byte int

	send_byte = 0
	recv_byte = 0

	var send_packet int
	var recv_packet int

	send_packet = 0
	recv_packet = 0

	fmt.Println(mac_address)
	handle, err = pcap.OpenOffline(pcapFile)
	fmt.Println(err)
	if err == nil {
		log.Fatal(err)
		defer handle.Close()

		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for {
			packet, err := packetSource.NextPacket()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Println("Error:", err)
				continue
			}

			ip_address, ports, send_byte_tip, recv_byte_tip, send_packet_tip, recv_packet_tip := PrintPacketInfo(control_ip, packet)

			send_byte += send_byte_tip
			recv_byte += recv_byte_tip

			send_packet += send_packet_tip
			recv_packet += recv_packet_tip

			if ip_address != "" {
				dst_ip = append(dst_ip, ip_address)
			}

			if ports != "" {
				sets = ip_address + "," + ports
				dst_ip_port = append(dst_ip_port, sets)
			}

		}

		if dst_ip_port != nil {
			unique_ip_dst = Unique_info(dst_ip)
			unique_ip_port_dst = Unique_info(dst_ip_port)

			fmt.Println(unique_ip_dst)

			fmt.Println("---------")
			ip_count = Count(unique_ip_dst, dst_ip)
			ip_port_count = Count(unique_ip_port_dst, dst_ip_port)

			fmt.Println(ip_port_count)

			// store execute
			Parse_port(ip_port_count)
			Parse_connect(control_ip, ip_count)
			Parse_traffic(control_ip, mac_address, send_byte, recv_byte, send_packet, recv_packet)
		}
	}
}
