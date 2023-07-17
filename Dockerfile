# Multi-stage build
# build stage
FROM golang:1.19-alpine AS build
WORKDIR /app

# Copy the Go module files to the working directory
# Download and cache Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Install curl    
# Dowload tailwind standalone client
# Give permission to tailwind executable
# Rename tailwind executable
RUN apk --no-cache add curl
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.3.3/tailwindcss-linux-x64 
RUN chmod +x tailwindcss-linux-x64 
RUN mv tailwindcss-linux-x64 tailwindcss

# Compile CSS
RUN ./tailwindcss -i ./resources/css/input.css -o ./dist/output.css

# Build the Go application
RUN go build -o server .

# run stage
FROM golang:1.19-alpine AS server

WORKDIR /app/

# Copy relevant files and folders from build stage
COPY --from=build /app/server ./
COPY --from=build /app/assets ./assets
COPY --from=build /app/dist ./dist
COPY --from=build /app/js ./js
COPY --from=build /app/templates ./templates

# Expose the port that the server listens on
EXPOSE 8080

# Set the entry point for the container
CMD ["./server"]