ARG UTASK_IMAGE=ovhcom/utask:v1.8.1

FROM ${UTASK_IMAGE} AS builder
WORKDIR /go/src/github.com/ovh/utask
COPY init init
COPY plugins plugins
COPY pkg /go/src/github.com/dsrhub/dsrhub/pkg
RUN make

FROM ${UTASK_IMAGE}
COPY templates  /app/templates
COPY --from=builder /go/src/github.com/ovh/utask/plugins /app/plugins
COPY --from=builder /go/src/github.com/ovh/utask/init    /app/init
COPY --from=builder /go/src/github.com/ovh/utask/utask   /app/utask
RUN chmod +x /app/utask
