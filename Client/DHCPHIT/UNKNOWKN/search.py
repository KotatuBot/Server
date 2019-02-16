import pings

class Search():
    def _init__(self):
        pass

    def response_machine(self,host):
        hit_host = []
        p = pings.Ping()
        for data in host:
            res = p.ping(data["IPAddress"])

            if res.is_reached():
                hit_host.append(data)

        return hit_host


