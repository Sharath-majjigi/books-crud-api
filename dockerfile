FROM golang:1.12.0-alpine3.9

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/Sharath-majjigi/books-crud-api/main
RUN cd /build && git clone https://github.com/Sharath-majjigi/books-crud-api.git

RUN cd /build/books-crud-api/main && go build

EXPOSE 8000

ENTRYPOINT [ "/build/books-crud-api/main/main" ]