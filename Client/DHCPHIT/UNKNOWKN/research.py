import subprocess
import shlex
import re

from scapy.all import *

class Research():
    def __init__(self):
        pass

    def genrate_mac(self,macaddress):
        mac_list = [macaddress[number:number+2] for number,j in enumerate(macaddress) if number % 2 ==0]
        return ":".join(mac_list)

    def search(self,ip):
        os_name = ""
        word = "^[0-9]*/"
        repattern = re.compile(word)

        os_word = "Aggressive OS guesses:*"
        repatternsd = re.compile(os_word)

        command = "nmap -O {0}".format(ip)
        try:
            data = subprocess.check_output(shlex.split(command))
            port_number = []
            status = []
            service = []
            datas = data.decode("utf-8").split("\n")
            for tips in datas:
                if len(repattern.findall(tips))!=0:
                    word = re.split(" +",tips)
                    port_number.append(word[0].split("/")[0])
                    status.append(word[1])
                    service.append(word[2])
                elif len(repatternsd.findall(tips))!=0:
                    os_listing = tips.split(":")
                    os_name = os_listing[1].split(" ")
                else:
                    pass

            Versions = []
            for ports in port_number:
                command = "nmap -sV -p {0} {1}".format(ports,ip)
                try:
                    data = subprocess.check_output(shlex.split(command),timeout=3)
                except:
                    print("timeout")
                dataed = data.decode("utf-8")
                test = dataed.split("\n")
                for tips in test:
                    if len(repattern.findall(tips))!=0:
                        wording = re.split(" +",tips)
                        try:
                            wordings = " ".join(wording[3:])
                            Versions.append(wordings)
                        except:
                            Versions.append('')
            if os_name != " ":
                os_name = "Unknown"

            if len(service) == 0:
                service = ["Unknown"]
               
            return port_number,status,service,Versions,os_name
        except:
            port_number = "Unknown"
            status = "Unknown"
            service = "Unknown"
            Versions = "Unknown"
            os_name = "Unknown"
            return port_number,status,service,Versions,os_name


    
    def main(self,ip):
        port,status,service,versions,os_name = self.search(ip)
        if type(port) == list:
            port = ",".join(port)
            service = ",".join(service)

        return port,status,service,versions,os_name


if __name__ == "__main__":
    research = Research()
    ip = "192.168.241.143"
    ports,status,services,versions,os_name  = research.main(ip)
    print(ports)
    print(services)
    print(versions)
    print(os_name)
