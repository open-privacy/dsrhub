# ARG UTASK_IMAGE=ovhcom/utask:v1.8.3
# TODO: upgrade to ovhcom/utask:v1.8.3 when it's ready
ARG UTASK_IMAGE=zhouzhuojie/utask:beta-with-middleware




FROM fufuhu/sql-migrate AS sql-migrate
FROM ${UTASK_IMAGE} AS builder
WORKDIR /go/src/github.com/ovh/utask
COPY init init
COPY plugins plugins
COPY pkg /go/src/github.com/dsrhub/dsrhub/pkg
RUN make




FROM frolvlad/alpine-glibc:alpine-3.10
RUN apk add --no-cache curl bash
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh
COPY --from=sql-migrate /bin/sql-migrate                     /usr/local/bin/sql-migrate
COPY --from=builder     /go/src/github.com/ovh/utask/plugins /app/plugins
COPY --from=builder     /go/src/github.com/ovh/utask/init    /app/init
COPY --from=builder     /go/src/github.com/ovh/utask/utask   /app/utask
COPY --from=builder     /go/src/github.com/ovh/utask/sql     /app/sql
COPY --from=builder     /app/static                          /app/static
COPY                    ./templates                          /app/templates
COPY                    ./dbconfig.yml                       /app/dbconfig.yml
RUN chmod +x /app/utask
EXPOSE 8081
WORKDIR /app
CMD ["/app/utask"]
