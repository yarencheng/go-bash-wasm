# Stage 1: Build the Go WASM binary
FROM golang:1.26-alpine AS go-builder

# Install git and binaryen (for wasm-opt)
RUN apk add --no-cache git binaryen

WORKDIR /src

# Clone the go-bash-wasm repository
RUN git clone --depth 1 https://github.com/yarencheng/go-bash-wasm.git .

# Build the WASM binary using the native Go build tool
# Targeting js/wasm for browser execution
RUN GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o main.wasm ./cmd/go-bash-wasm/

# Optimize the WASM binary, TODO: add --minify-imports-and-exports
RUN wasm-opt -O3 \
    --enable-bulk-memory \
    --enable-nontrapping-float-to-int \
    --enable-sign-ext \
    --strip-producers \
    --strip \
    --strip-producers \
    --flatten \
    --coalesce-locals \
    --simplify-locals-notee \
    --inlining-optimizing \
    main.wasm -o main_fast.wasm \
    && mv main_fast.wasm main.wasm

# Find the wasm_exec.js file and copy it to the build directory
RUN find / -name wasm_exec.js -exec cp {} . \;

# Stage 2: Build the Svelte application
FROM node:20-alpine AS builder

WORKDIR /app

# Copy dependency manifests
COPY package*.json ./

# Install dependencies (only devDependencies are needed for build)
RUN npm ci

# Copy source code
COPY . .

# Run unit tests
RUN npm run test

# Build the application
RUN npm run build

# Stage 3: Serve with Nginx
FROM nginx:stable-alpine

# Copy the build output from the builder stage
COPY --from=builder /app/build /usr/share/nginx/html

# Copy the WASM artifacts from the go-builder stage
COPY --from=go-builder /src/main.wasm /usr/share/nginx/html/main.wasm
COPY --from=go-builder /src/wasm_exec.js /usr/share/nginx/html/

# Copy custom nginx config
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

