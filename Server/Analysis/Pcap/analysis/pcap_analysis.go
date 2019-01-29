package analysis

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"strings"
)

func PrintPacketInfo(control_ip string, packet gopacket.Packet) (string, string, int, int, int, int) {

	var dst_port string
	var str_remoteaddr string
	var dst_ip string
	var connect_ip string

	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	ip, _ := ipLayer.(*layers.IPv4)

	str_remoteaddr = ip.SrcIP.String()
	dst_ip = ip.DstIP.String()

	send_byte, recv_byte, send_packet, recv_packet := Get_Byte(control_ip, str_remoteaddr, dst_ip, packet)
	if tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		if str_remoteaddr == control_ip {
			dst_ip := ip.DstIP.String()
			dst_port = strings.Split(tcp.DstPort.String(), "(")[0]
			connect_ip = dst_ip
		}
	} else {
		udpLayer := packet.Layer(layers.LayerTypeUDP)
		if udpLayer != nil {
			udp, _ := udpLayer.(*layers.UDP)
			if str_remoteaddr == control_ip {
				dst_ip := ip.DstIP.String()
				dst_port = strings.Split(udp.DstPort.String(), "(")[0]
				connect_ip = dst_ip
			}
		}

	}

	return connect_ip, dst_port, send_byte, recv_byte, send_packet, recv_packet

}
