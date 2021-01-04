FROM golang:1.14.4-stretch

WORKDIR /usr/local/go/src/amfLoadBalancer
EXPOSE 40000 1541

COPY go.mod .

RUN go mod download

RUN rm -rf /usr/local/go/src/vendor/golang.org/x/net 

ADD lib ./lib
ADD src ./src
VOLUME ["/usr/local/go/src/amfLoadBalancer/config"]

RUN go build -a --installsuffix nocgo -o bin/amfLoadBalancer src/main.go