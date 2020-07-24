# ARG UTASK_IMAGE=ovhcom/utask:v1.8.3
# TODO: upgrade to ovhcom/utask:v1.8.3 when it's ready
ARG UTASK_IMAGE=zhouzhuojie/utask:beta-with-middleware




FROM ${UTASK_IMAGE} AS builder
WORKDIR /go/src/github.com/ovh/utask
COPY init init
COPY plugins plugins
COPY pkg /go/src/github.com/dsrhub/dsrhub/pkg
RUN make




FROM ${UTASK_IMAGE}
COPY --from=builder /go/src/github.com/ovh/utask/plugins /app/plugins
COPY --from=builder /go/src/github.com/ovh/utask/init    /app/init
COPY --from=builder /go/src/github.com/ovh/utask/utask   /app/utask
COPY --from=builder /go/src/github.com/ovh/utask/sql     /app/sql
RUN go get -v github.com/rubenv/sql-migrate/...
COPY ./templates    /app/templates
COPY ./dbconfig.yml /app/dbconfig.yml
RUN chmod +x /app/utask
WORKDIR /app
