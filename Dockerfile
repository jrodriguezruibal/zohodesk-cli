FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /root/

COPY zohodesk-cli /usr/local/bin/zohodesk-cli

RUN mkdir -p /root/.config/zohodesk-cli

ENTRYPOINT ["zohodesk-cli"]
CMD ["version"]