# docker build -t [image_name]:[tag] .
FROM golang:1.11.6-alpine as builder
WORKDIR /go/src/
ENV GO111MODULE=on

# Copy vendor first and build for cache layers
COPY go.mod .
COPY go.sum .
COPY vendor vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -mod=vendor -ldflags="-w -s" -v ./vendor/...

# Copy source codes and build
COPY . .
ARG GITVERSION
ARG GITBRANCH
ARG GITREVISION
ENV PKGPATH=gitlab.com/rayshih/go-api-template/gpt
RUN export BUILDDATE=$(date +"%Y-%m-%dT%H:%M:%SZ") && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -mod=vendor -ldflags="-w -s -X $PKGPATH.gitCommit=$GITREVISION -X $PKGPATH.appVersion=$GITVERSION -X $PKGPATH.buildDate=$BUILDDATE -X $PKGPATH.gitBranch=$GITBRANCH" -v ./...

# For minimal image
FROM alpine:3.9
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/go-api-template .
COPY local.toml .
CMD ./go-api-template -config local.toml
