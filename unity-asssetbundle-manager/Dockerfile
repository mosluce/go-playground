FROM golang:1.9

WORKDIR /go/src/app
COPY . .

VOLUME /go/src/app/storage

ENV PORT=80
ENV GIN_MODE=release

EXPOSE 80

RUN go-wrapper download   
RUN go-wrapper install    

CMD ["go-wrapper", "run"] 