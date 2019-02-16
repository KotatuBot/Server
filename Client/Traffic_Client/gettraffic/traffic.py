import subprocess
import shlex
import re
import time
import datetime
import json
import os

class Get_trafic():
    def __init__(self):
        self.file_name = "/opt/Data/write.json" 
        self.interface = ".*eth0*"

    def main(self):
        command = "cat /proc/net/dev"
        #pattern = ".*eth0*"
        pattern = self.interface
        patterns = re.compile(pattern)
        data = subprocess.check_output(shlex.split(command))

        for tip in data.split('\n'):
            if patterns.match(tip):
                split_data = tip.split(' ')
                interface_list = [d for d in split_data if d != '']
                break
        date = datetime.datetime.now()
        date_format = "{0}-{1}-{2} {3}:{4}:{5}".format(date.year,date.month,date.day,date.hour,date.minute,date.second)
        trafic_data = {
            "Date":date_format,
            "Recive_byte" : interface_list[1],
            "Recive_packet" :interface_list[2],
            "Send_byte": interface_list[9],
            "Send_packet": interface_list[10],
        }

        with open(file_name,"w") as fd:
            json.dump(trafic_data,fd)
            os.chmod(file_name,777)


        return "success"

if __name__ == "__main__":
   gt =  Get_trafic()
   string_concat = gt.main()
   print(string_concat)
