# This image is a microservice in golang for the Degree chaincode
FROM golang:1.14.6-alpine AS build

COPY ./ /go/src/github.com/channel1-chaincode-on-fabric
WORKDIR /go/src/github.com/channel1-chaincode-on-fabric

# Build application
RUN sh build.sh

# Production ready image
# Pass the binary to the prod image
FROM alpine:3.11 as prod

COPY --from=build /go/src/github.com/channel1-chaincode-on-fabric/build/chaincode /app/chaincode


USER 1000

WORKDIR /app
# RUN chmod 777 ./chaincode
CMD [ "chaincode" ]
