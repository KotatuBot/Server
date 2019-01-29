FROM golang
RUN apt-get update 
RUN apt-get upgrade -y
RUN apt-get install -y vim
RUN apt-get install -y libpcap-dev
RUN mkdir App
WORKDIR App
RUN mkdir Data
RUN mkdir analysis 
ADD analysis analysis
ADD main.go .
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/oschwald/geoip2-golang
RUN go get github.com/google/gopacket
RUN go build
CMD ["sleep","80000"]
