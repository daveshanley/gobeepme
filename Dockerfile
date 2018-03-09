FROM golang:1.8

RUN mkdir -p /go/src/github.com/daveshanley/gobeepme
WORKDIR /go/src/github.com/daveshanley/gobeepme

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/daveshanley/gobeepme
COPY . /go/src/github.com/daveshanley/gobeepme

RUN go get -d -v ./...
RUN go install -v ./...

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/gobeepme -service

# Document that the service listens on port 9443
EXPOSE 9443
