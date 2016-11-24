FROM alpine
MAINTAINER jspc <james@zero-internet.org.uk>

ADD ghost-post-generator /ghost-post-generator

ENTRYPOINT ["/ghost-post-generator"]
