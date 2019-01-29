iptables -t nat -A POSTROUTING -o enx3495db11dfe2 -j MASQUERADE
iptables -A FORWARD -i enx3495db11dfe2 -o wlp3s0 -m state --state RELATED,ESTABLISHED -j ACCEPT
iptables -A FORWARD -i wlp3s0 -o enx3495db11dfe2 -j ACCEPT
