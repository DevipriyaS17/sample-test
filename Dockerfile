FROM alpine:latest

# Install CA certificates for HTTPS connections
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy the compiled binary
COPY publish_output/hello-web-server /app/

# Set execute permissions on the binary
RUN chmod +x /app/hello-web-server

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["/app/hello-web-server"]
