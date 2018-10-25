FROM golang as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build

RUN go build -a -installsuffix cgo -ldflags '-extldflags "-s -L./indigo/lib -Wl,-rpath=./indigo/lib -lindigo"' -o service http-server-indigo.go
#FROM debian:stable-slim
#FROM bitnami/minideb:jessie
FROM busybox:glibc
COPY --from=builder /build/service /app/
COPY --from=builder /build/indigo/lib/libindigo.so /app/indigo/lib/

#COPY --from=builder /lib/x86_64-linux-gnu/libdl-2.24.so /lib/libdl-2.24.so
#COPY --from=builder /lib/x86_64-linux-gnu/libdl.so.2 /lib/libdl.so.2
COPY --from=builder /lib/x86_64-linux-gnu/libgcc_s.so.1 /lib/libgcc_s.so.1

WORKDIR /app
CMD ["./service"]
