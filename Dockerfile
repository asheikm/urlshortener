# Use centos 7 Base image
FROM centos:centos7

# Set evn variables 

# App listen on port 8080
ENV LISTEN_PORT=8080

# App log file path
ENV LOGFILE_PATH="./log/urlshortener_log.out"

# App short domain name
ENV SHORT_DOMAIN="localhost:8080/"

# Directory for binary and log
RUN mkdir -p /urlshortener/log

# Work directory
WORKDIR /urlshortener

# Copy binary to workdir
COPY bin/urlshortener .

# Expose port outside world 
EXPOSE 8080

# Run binary from command line
CMD ["./urlshortener"]
