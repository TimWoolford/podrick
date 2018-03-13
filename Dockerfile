FROM gcr.io/distroless/base

COPY template    /template
COPY static      /static
COPY config      /config
COPY bin/podrick /podrick

ENTRYPOINT ["/podrick"]