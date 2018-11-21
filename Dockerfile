FROM golang:1.11.2

RUN wget https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && mv dep-linux-amd64 /usr/bin/dep && chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/github.com/PuKoren/convenient-api

COPY dbs/*.txt dbs/
COPY dbs/*.mmdb dbs/

COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only

COPY . .

RUN go build main.go

ENTRYPOINT ./main
