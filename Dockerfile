FROM golang:1.18 as builder
RUN mkdir /build
ADD go.mod go.sum main.go /build/
WORKDIR /build
RUN export CGO_ENABLED=0 && go build -o helloapi

FROM alpine
RUN adduser -S -D -H -h /app app
USER app
COPY --from=builder /build/helloapi /app/
WORKDIR /app
#ENV PORT=${LISTENPORT:-8083}
#EXPOSE ${LISTENPORT}
CMD ["./helloapi"]
