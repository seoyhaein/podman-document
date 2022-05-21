FROM alpine:latest
RUN apk update && apk add --no-cache bash nano
ADD ../shellscripts/procConfirm.sh /usr/local/bin/procConfirm.sh
ENTRYPOINT ["/usr/local/bin/procConfirm.sh"]
