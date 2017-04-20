FROM busybox
MAINTAINER benton@bentonroberts.com

ADD pkg/htest /htest
EXPOSE 80
ENTRYPOINT ["/htest"]
