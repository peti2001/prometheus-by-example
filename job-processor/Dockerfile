FROM alpine:3.2

RUN apk --update upgrade && \
    apk add curl ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

ADD job-processor /job-processor

CMD /job-processor