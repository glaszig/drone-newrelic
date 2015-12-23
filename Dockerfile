# Docker image for the Drone New Relic plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-newrelic
#     make deps build docker

FROM alpine:3.2
ADD drone-newrelic /bin/
ENTRYPOINT ["/bin/drone-newrelic"]
