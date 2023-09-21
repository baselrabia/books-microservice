# FROM golang:1.21

# # Set the working directory inside the container
# WORKDIR /app

# # Copy the Go application code into the container
# COPY . .

# # Build the Go application
# RUN go build -o main .

# # Expose the port the application will run on
# EXPOSE 8080

# # Command to run the executable
# CMD ["./main"]

 

FROM golang:1.21.0 as builder
WORKDIR /app
RUN go mod init github.com/baselrabia/book-microservice
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /book-microservice

FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=builder /book-microservice /book-microservice
ENV PORT 8080
USER nonroot:nonroot
CMD ["/book-microservice"]