FROM golang
RUN apt-get update
RUN apt-get -y upgrade
RUN apt-get install -y vim
RUN apt-get install -y python-pip
RUN pip install pyftpdlib
RUN mkdir App
WORKDIR App
RUN mkdir Data
ADD server.py .
ADD start.sh .
RUN chmod +x start.sh
EXPOSE 21
CMD ["sh","start.sh"]
