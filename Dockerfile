FROM creativebit/go-dep

LABEL maintainer="hi@jhsc.io"

WORKDIR /go/src/github.com/jhsc/rtsupport

COPY . .
