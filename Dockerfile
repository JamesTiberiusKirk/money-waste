FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' money-waste .

FROM alpine
COPY --from=builder /build/money-waste /app/
WORKDIR /app
CMD ["./money-waste"]
