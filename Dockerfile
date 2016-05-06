FROM haproxy:1.6-alpine
MAINTAINER 	Viktor Farcic <viktor@farcic.com>

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY go-demo /usr/local/bin/go-demo
RUN chmod +x /usr/local/bin/go-demo

EXPOSE 8080

CMD ["go-demo"]
