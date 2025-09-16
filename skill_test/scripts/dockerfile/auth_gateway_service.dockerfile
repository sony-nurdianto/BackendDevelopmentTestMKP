FROM golang:1.24.4-bookworm

COPY service/gRPC/auth/build/auth_gateway_service  /usr/bin/auth_gateway_service

EXPOSE 3000

CMD ["auth_gateway_service"]
