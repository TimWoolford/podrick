FROM gcr.io/distroless/base

ADD  template    /template
ADD  static      /static
COPY bin/podrick /podrick

ENTRYPOINT ["/podrick"]