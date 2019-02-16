cd ./Pcap_Client
./Pcap_Client &

cd ../

cd ./Traffic_Client
sudo ./Traffic_Client &

cd ../
cd ./DHCPHIT/UNKNOWKN
sudo python3 main.py &

sleep 80000000

