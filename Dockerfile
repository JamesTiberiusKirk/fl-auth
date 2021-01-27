# FROM golang:1.15-alpine as builder 

# LABEL maintainer="Dumitru Vulpe <dumitru.v.dv@gmail.com>"

# RUN mkdir /app
# WORKDIR /app

# COPY go.mod go.sum ./

# RUN go mod download

# COPY . .

# RUN go build fl-atuh

# FROM alpine
# RUN mkdir /app
# WORKDIR /app
# COPY --from=builder /app/fl-auth /app

# EXPOSE 3000

# CMD ["./fl-auth"]




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