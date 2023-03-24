# Telling to use Docker's golang ready image
FROM golang:1.12.0-alpine3.9

# Create app folder 
RUN mkdir /app
# Copy our file in the host contianer to our contianer
ADD . /app
# Set /app to the go folder as workdir
WORKDIR /app
# Generate binary file from our /app
RUN go build 
# Expose the port 3000
EXPOSE 8000:8000
# Run the app binarry file 
CMD ["/app/main"]