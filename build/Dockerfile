FROM golang:1.17-alpine AS build

RUN apk update && apk add --no-cache git

# Create a user group 'theater'
# Create a user 'usher' under 'theater'
RUN addgroup theater && adduser --disabled-password -DH -G theater usher

# Set the Current Working Directory inside the container
WORKDIR /tmp/usher

ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./app/usher ./cmd/usher

# Start fresh from a smaller image
FROM alpine:latest

ENV PORT 8081
EXPOSE 8081

# Import users and groups
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group

# Chown all the files to the app user.
RUN mkdir /app
RUN chown -R usher:theater /app

# Switch user usher
USER usher:theater

COPY --from=build /tmp/usher/app /app

# Run the binary program produced by `go install`
CMD ["/app/usher"]