# First stage: build the executable.
FROM golang:1.15.8-buster as builder

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/emaele/rss-telegram-notifier

# Import the code from the context.
COPY .  .

# Build the executable
RUN go install

# Final stage: the running container.
FROM alpine:3.13.1

# Copy the built executable
COPY --from=builder /go/bin/rss-telegram-notifier .

# Expose the port
EXPOSE 26009

# Run the compiled binary.
CMD ["./rss-telegram-notifier"]