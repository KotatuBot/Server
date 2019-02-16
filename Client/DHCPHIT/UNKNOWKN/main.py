from research import Research
from use_api import Use_Api
from get_lease import Get_Lease4
from search import Search
import time

# leaseチェック
gl = Get_Lease4()

while True:
    lease4_dict = gl.operate_lease4()
    print("sleep")
    time.sleep(5)
    research = Research()
    api = Use_Api("http://127.0.0.1:8000/api/control/","http://127.0.0.1:8000/api/unknown/")

    for lease4 in lease4_dict:
        ip = lease4["IPAddress"]
        # IPチェック(ping処理)
        macaddress = research.genrate_mac(lease4["MACaddress"])
        status = api.confirmation(ip,macaddress)
        status_ago = api.unknow_confirmation(ip,macaddress)
        print(status_ago)
        # 登録機器に存在しないならば
        if status == "No":
            # 過去に登録されていないか
            if status_ago == "No":
                ports,status,services,versions,os_name  = research.main(ip)
                manufacture = "Unknown"
                print("Register")
                print("IP:"+ip)
                print("MacAddress:"+macaddress)
                print("OS:"+os_name)
                print("port:"+ports)
                print("Service:"+services)
                print("manufacture:"+manufacture)
                api.posts(ip,macaddress,os_name,ports,services,manufacture)
