#using python3
import ipaddress
import pymysql.cursors

from search import Search

class Get_Lease4():
   def __init__(self):
       self.connection = ""

   def db_connect(self):
       self.connection = pymysql.connect(host='localhost',
                            user='root',
                            password='dhcpdb',
                            db='dhcpdb',
                            charset='utf8',
                            cursorclass=pymysql.cursors.DictCursor)

   def all_leaseas(self):
       lease4_data = []
       with self.connection.cursor() as cursor:
           sql = "SELECT * FROM lease4"
           cursor.execute(sql)
           results = cursor.fetchall()
           for data in results:
               lease4_data.append(data)
           return lease4_data

   def close(self):
       self.connection.close()


   def operate_lease4(self):
       self.db_connect()
       lease4 = self.all_leaseas()
       self.close()
       get_data = self.get_info(lease4)
       return get_data


   def get_info(self,lease4):
       lease4_data = []
       ip_list = []
       for l4 in lease4:
           ipv4_class = ipaddress.IPv4Address(l4["address"])
           ipaddress4 = str(ipv4_class)
           hwaddress = l4["hwaddr"].hex()
           expire_time = "{0:%Y-%m-%d %H:%M:%s}".format(l4["expire"]) 
           dhcp_data = {"IPAddress":ipaddress4,"MACaddress":hwaddress,"Exper":expire_time}
           lease4_data.append(dhcp_data)
       return lease4_data

if __name__ == "__main__":
  gl = Get_Lease4()
  se = Search()
  lease4_data = gl.operate_lease4()
  r = se.response_machine(lease4_data)
  print(r)
