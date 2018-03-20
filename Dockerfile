FROM gcr.io/distroless/base

COPY web/template /template
COPY web/static   /static
COPY config       /config
COPY bin/podrick  /podrick

ENTRYPOINT ["/podrick"]