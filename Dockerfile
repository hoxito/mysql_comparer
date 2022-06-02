# Docker para desarrollo
FROM golang:1.14.3-buster

WORKDIR /go/src/mysqlbinlogparser

# Puerto de stats Service y debug
EXPOSE 3010

CMD ["go" , "run" , "/go/src/mysqlbinlogparser"]