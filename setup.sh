mkdir /opt/Data
chmod 777 /opt/Data
mkdir /opt/Split
chmod 777 /opt/Split
cd $HOME
mkdir .go
echo 'export GOPATH=$HOME/.go' >> .zshrc

# DHCP HIT
sudo apt-get install nmap
sudo apt-get install wireshark
sudo apt-get isntall tshark
pip3 install ipaddress
pip3 install pymysql
pip3 install requests

# Clinet and Server
go get github.com/go-sql-driver/mysql
go get github.com/oschwald/geoip2-golang
go get github.com/google/gopacket

