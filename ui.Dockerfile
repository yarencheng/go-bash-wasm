# Stage 1: Build and Test
FROM golang:1.26-alpine AS go-builder

# Install binaryen (for wasm-opt)
ARG BINARYEN_VERSION=123
RUN apk add --no-cache curl tar ca-certificates libstdc++ \
    && BINARYEN_ARCH=$(uname -m | sed 's/arm64/aarch64/') \
    && curl -L "https://github.com/WebAssembly/binaryen/releases/download/version_${BINARYEN_VERSION}/binaryen-version_${BINARYEN_VERSION}-${BINARYEN_ARCH}-linux.tar.gz" | tar -xz \
    && mv binaryen-version_${BINARYEN_VERSION}/bin/wasm-opt /usr/local/bin/ \
    && rm -rf binaryen-version_${BINARYEN_VERSION} \
    && apk del curl tar

WORKDIR /src

# Copy go.mod and go.sum (if present) for caching dependencies
COPY go.mod go.sum* ./
RUN go mod download

# Copy project files
COPY . .

# Run all unit tests
RUN go test -v ./...

# Build the WASM binary using the native Go build tool
# Targeting js/wasm for browser execution
RUN GOOS=js GOARCH=wasm go build -ldflags="-s -w -X 'github.com/yarencheng/go-bash-wasm/internal/app.MachType=wasm32-unknown-wasi'" -o main.wasm ./cmd/go-bash-wasm/

ARG OPTIMIZE=fast
RUN if [ "${OPTIMIZE}" = "fast" ]; then \
    wasm-opt -O4 \
    --enable-bulk-memory \
    --enable-nontrapping-float-to-int \
    --enable-sign-ext \
    --enable-simd \
    --enable-tail-call \
    --enable-threads \
    --inlining \
    --inlining-optimizing \
    --precompute \
    --precompute-propagate \
    --gufa \
    --simplify-locals \
    --coalesce-locals \
    --reorder-locals \
    --dce \
    --dae \
    main.wasm -o main_fast.wasm \
    && mv main_fast.wasm main.wasm; \
    fi

RUN if [ "${OPTIMIZE}" = "small" ]; then \
    wasm-opt -Oz \
    --enable-bulk-memory \
    --enable-nontrapping-float-to-int \
    --enable-sign-ext \
    --strip-debug \
    --strip-dwarf \
    --strip-producers \
    --converge \
    --dce \
    --remove-unused-module-elements \
    --duplicate-function-elimination \
    --simplify-locals \
    --coalesce-locals \
    --reorder-locals \
    --vacuum \
    --memory-packing \
    main.wasm -o main_small.wasm \
    && mv main_small.wasm main.wasm; \
    fi

# Stage 2: Build the Svelte application
FROM node:20-alpine AS ui-builder

WORKDIR /app

# Copy dependency manifests
COPY ui/package*.json ./

# Install dependencies (only devDependencies are needed for build)
RUN npm ci

# Copy source code
COPY ui/ .

# Run unit tests
RUN npm run test

# Update sitemap.xml lastmod with current date
RUN sed -i "s|<lastmod>.*</lastmod>|<lastmod>$(date +%Y-%m-%d)</lastmod>|g" static/sitemap.xml

# Build the application
RUN npm run build

# Stage 3: Serve with Nginx
FROM nginx:stable-alpine

# Copy the build output from the builder stage
COPY --from=ui-builder /app/build /usr/share/nginx/html

# Copy the WASM artifacts from the go-builder stage
COPY --from=go-builder /src/main.wasm /usr/share/nginx/html/main.wasm
COPY --from=go-builder /usr/local/go/lib/wasm/wasm_exec.js /usr/share/nginx/html/

# Copy custom nginx config
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

