FROM golang:1.13

# Set GOPATH/GOROOT environment variables
# RUN mkdir -p /go/src
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

RUN mkdir -p /app
# go get all of the dependencies
# RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Set up app
# WORKDIR /go/src/github.com/hadv/eb-echo-docker
ADD ./config /config

# RUN go get -u github.com/labstack/echo/...
# RUN go get github.com/cortesi/modd/cmd/modd
RUN env GO111MODULE=on go get github.com/cortesi/modd/cmd/modd

# RUN go build -v

EXPOSE 3000
RUN chmod +x /config/run.sh

CMD ["/config/run.sh"]