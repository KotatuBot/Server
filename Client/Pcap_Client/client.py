import os
import sys

from ftplib import FTP

class Client():
    def __init__(self,IP):
        self.IP = IP

    def delete_file(self,path):
        os.remove(path)

    def send_Pcap(self,ftp):
        """
        Send Pcap data
        """
        current = os.getcwd()
        Dir_name = "http.pcapng"
        os.chdir("Data")
        with open(Dir_name, "rb") as f:
            ftp.storlines("STOR "+Dir_name, f)
        self.delete_file(Dir_name)
        os.chdir(current)

    def send_Trafic(self,ftp):
        current = os.getcwd()
        Dir_name = "write.json"
        os.chdir("Data")
        with open(Dir_name, "rb") as f:
            ftp.storlines("STOR "+Dir_name, f)
        self.delete_file(Dir_name)
        os.chdir(current)

    def pcap_main(self):
        ftp = FTP(
                    self.IP,
                    "user",
                    passwd="password"
                            )

        self.send_Pcap(ftp)
        ftp.quit();

    def traffic_main(self):
        ftp = FTP(
                    self.IP,
                    "user",
                    passwd="password"
                            )

        self.send_Trafic(ftp)
        ftp.quit();

    def main(self,data):
        if data == "pcap":
            self.pcap_main()
        elif data == "traffic":
            self.traffic_main()
        else:
            pass

if __name__ == "__main__":
    # python client.py 127.0.0.1 pcap|traffic
    args = sys.argv
    print(args)
    client = Client(args[1])
    client.main(args[2])
    print("test")
