FROM gcr.io/distroless/base

ADD template /template
COPY bin/podrick /podrick

ENTRYPOINT ["/podrick"]