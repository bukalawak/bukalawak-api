# syntax=docker/dockerfile:1
   
FROM alpine:3.17.2

RUN apk add --no-cache musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

WORKDIR /
COPY . .
RUN go build -o /bukalawak-api
CMD ["/bukalawak-api"]
EXPOSE 8000