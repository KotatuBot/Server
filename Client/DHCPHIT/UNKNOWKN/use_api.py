import requests
import json

class Use_Api():
    def __init__(self,confirmation_url,post_url):
        # /api/control
        self.confirmation_url = confirmation_url
        # /api/unkown
        self.post_url = post_url
        self.check_url = post_url

    def confirmation(self,ip,macaddress):
        """
        登録されているかを確認する関数
        """
        status = ""
        payload = {'Control_ip':ip, 'MacAddress':macaddress}
        r = requests.get(self.confirmation_url, params=payload)
        # データが存在する
        if len(r.text) != 2:
            status = "Yes"
        # データが存在しない
        else:
            status = "No"
        return status

    def fileterconfirmation(self,ip,macaddress):
        """
        登録されているかを確認する関数
        """
        unknown_url = "http://127.0.0.1:8000/api/unknownfilter/"
        status = ""
        payload = {'Unknown_ip':ip, 'MacAddress':macaddress}
        r = requests.get(unknown_url, params=payload)
        # データが存在する
        if len(r.text) != 2:
            status = "Yes"
        # データが存在しない
        else:
            status = "No"
        return status


    def unknow_confirmation(self,ip,macaddress):
        """
        Unkownに登録されているかを確認する関数
        """
        status = ""
        payload = {'Control_ip':ip, 'MacAddress':macaddress}
        r = requests.get(self.post_url, params=payload)
        test = r.text
        # データが存在する
        if len(r.text) != 2:
            status2 = self.fileterconfirmation(ip,macaddress)
            # unknow
            if status2 == "YES":
                status = "No"
            else:
                status = "Yes"
        # データが存在しない
        else:
            status = "No"
        return status

    def posts(self,ip,macaddress,os,port,service,manufacture):
        """
        未登録APIに値を追加する.
        """
        payload = {'Unknown_ip':ip,'MacAddress':macaddress,'OS':os,'Port':port,'Service':service,'Manufacture':manufacture}
        headers = {'Content-Type':'application/json'}
        r = requests.post(self.post_url,data=json.dumps(payload),headers=headers)


if __name__ == "__main__":
    ua = Use_Api("http://127.0.0.1:8000/api/control/","http://127.0.0.1:8000/api/unknown/")
    ua.posts("t","t","t","t","t","t")
