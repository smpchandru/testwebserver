FROM golang:alpine as builder
RUN mkdir -p /build/tmp 
WORKDIR /build
RUN wget https://github.com/smpchandru/testwebserver/archive/master.zip && \
unzip master.zip
WORKDIR testwebserver-master
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/testwebserver-master/main /app/
WORKDIR /app
CMD ["./main"]
