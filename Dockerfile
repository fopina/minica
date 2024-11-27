FROM alpine:3.20 AS alpine

COPY minica /usr/bin/minica
ENTRYPOINT ["/usr/bin/minica"]

FROM scratch

COPY minica /usr/bin/minica
ENTRYPOINT ["/usr/bin/minica"]
