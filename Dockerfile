FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Sử dụng ARG để truyền biến môi trường từ build stage sang stage cuối
ARG JWT_SIGNER_KEY
ARG DB_HOST
ARG DB_PORT
ARG DB_USER
ARG DB_PASSWORD
ARG DB_NAME

# Build ứng dụng
RUN go build -o main

# Stage cuối cùng
FROM gcr.io/distroless/base-debian11

WORKDIR /app

# Copy binary từ builder stage
COPY --from=builder /build/main .

# Set biến môi trường trong stage cuối. Giá trị mặc định là chuỗi rỗng.
ENV JWT_SIGNER_KEY=${JWT_SIGNER_KEY}
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_USER=${DB_USER}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_NAME=${DB_NAME}

EXPOSE 8080

CMD ["./main"]