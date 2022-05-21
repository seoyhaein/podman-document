FROM alpine:latest
RUN apk update && apk add --no-cache bash nano
RUN mkdir -p /opt/data
RUN chmod a+rw /opt/data
COPY ./usage.md /opt/data