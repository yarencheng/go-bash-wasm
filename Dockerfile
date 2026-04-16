# Stage 1: Build and Test
FROM golang:1.26-alpine AS builder

WORKDIR /src

# Copy go.mod and go.sum (if present) for caching dependencies
COPY go.mod go.sum* ./
RUN go mod download

# Copy project files
COPY . .

# 1. Run all unit tests
RUN go test -v ./...

# 2. Build web assembly output
# Change from wasip1/wasm to js/wasm
RUN GOOS=js GOARCH=wasm go build -o /out/main.wasm ./cmd/go-bash-wasm/
RUN cp $(go env GOROOT)/lib/wasm/wasm_exec.js /out/wasm_exec.js
COPY index.html /out/index.html

# Stage 2: Nginx Runner
FROM nginx:alpine

# Copy built artifacts from the builder stage
COPY --from=builder /out/main.wasm /usr/share/nginx/html/main.wasm
COPY --from=builder /out/wasm_exec.js /usr/share/nginx/html/wasm_exec.js
COPY --from=builder /out/index.html /usr/share/nginx/html/index.html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
