FROM alpine:3.12.3 as kctl
RUN apk add curl
RUN curl -L "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl" \
      -o /kubectl
RUN chmod +x /kubectl

FROM gcr.io/gcp-runtimes/go1-builder:1.13 as builder
ENV CGO_ENABLED=0
WORKDIR /go/src/
COPY image/go.mod .
RUN /usr/local/go/bin/go mod download
COPY main.go .
RUN /usr/local/go/bin/go build -v -o /usr/local/bin/config-function ./

FROM alpine:3.12.3 as calicoctl
RUN apk add curl
RUN curl -L  "https://github.com/projectcalico/calicoctl/releases/download/v3.18.1/calicoctl" \
      -o /calicoctl
RUN chmod +x /calicoctl

FROM alpine:3.12.3
RUN apk update \
  && apk add bash \
  && rm -rf /var/cache/apk/*
COPY --from=kctl /kubectl /usr/local/bin/kubectl
COPY --from=calicoctl /calicoctl /usr/local/bin/calicoctl
COPY --from=builder /usr/local/bin/config-function /usr/local/bin/config-function
CMD ["/usr/local/bin/config-function"]
