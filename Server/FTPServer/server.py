import pyftpdlib.authorizers
import pyftpdlib.handlers
import pyftpdlib.servers
import os

srvRoot = os.getcwd()+"/Data/";
print(srvRoot)
auth = pyftpdlib.authorizers.DummyAuthorizer()
auth.add_user('user', 'password', srvRoot, perm='elradfmw')

hndler = pyftpdlib.handlers.FTPHandler
hndler.authorizer = auth

server = pyftpdlib.servers.FTPServer(("127.0.0.1", 21), hndler)

try:
        server.serve_forever();
except KeyboardInterrupt:
        print('interrupted!')

