FROM golang
RUN apt-get update 
RUN apt-get upgrade -y
RUN apt-get install -y vim
RUN mkdir App
WORKDIR App
RUN mkdir Data
RUN mkdir exectraffic
ADD exectraffic exectraffic
ADD main.go .
RUN go get github.com/go-sql-driver/mysql
RUN go build
CMD ["sleep","80000"]
