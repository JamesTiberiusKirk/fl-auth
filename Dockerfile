FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' fl-auth .

FROM scratch
COPY --from=builder /build/fl-auth /app/
WORKDIR /app
EXPOSE 3000
CMD ["./fl-auth"]