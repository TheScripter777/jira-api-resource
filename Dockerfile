FROM golang:1.12 AS builder

LABEL maintainer="Guillaume Rivest" \
    authors="Guillaume Rivest" \
    version="0.0.1" \
    description="Jira API resource for Concourse CI"

# Copy everything from the jira-api-ressource module to /app in the image
COPY . /app

# Set the working directory to /app where our code and scripts are
WORKDIR /app

# Environment variables
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

# Set the go module file location
ENV GOMOD /app/go.mod

# Launch the make tool on the default target
RUN make

FROM alpine:edge AS resource

RUN apk --no-cache add \
        curl \
        jq \
        bash \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true \
;

# Copy the built binary into the bin folder
COPY --from=builder /app/bin/jiraApiResource /usr/local/bin/

# Copy assets
COPY assets/check /opt/resource/check
COPY assets/in /opt/resource/in
COPY assets/out /opt/resource/out

RUN chmod +x /opt/resource/check /opt/resource/in /opt/resource/out

WORKDIR /tmp/build

FROM resource
