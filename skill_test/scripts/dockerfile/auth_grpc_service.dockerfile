FROM golang:1.24.4-bookworm

COPY service/gRPC/auth/build/auth_grpc_service  /usr/bin/auth_grpc_service

EXPOSE 50051

CMD ["auth_grpc_service"]
