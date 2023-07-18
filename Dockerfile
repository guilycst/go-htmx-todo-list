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
RUN ./tailwindcss -i ./internal/web/src/input.css -o ./dist/output.css

# Copy js files to dist
RUN cp -R ./internal/web/src/*.js ./dist/

# Build the Go application
RUN go build -o server ./cmd/server

# run stage
FROM golang:1.19-alpine AS server

WORKDIR /app/

# Copy relevant files and folders from build stage
COPY --from=build /app/server ./
COPY --from=build /app/internal/web/templates ./templates
COPY --from=build /app/dist ./dist
COPY --from=build /app/internal/web/public ./public

#Set env
ENV TEMPLATES_DIR=/app/templates
ENV DIST_DIR=/app/dist
ENV PUB_DIR=/app/public

# Expose the port that the server listens on
EXPOSE 8080

# Set the entry point for the container
CMD ["./server"]