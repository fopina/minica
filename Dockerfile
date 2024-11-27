FROM scratch

COPY minica /usr/bin/minica
ENTRYPOINT ["/usr/bin/minica"]
