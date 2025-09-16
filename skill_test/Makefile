gen_auth_proto:	
	 cd proto && buf generate --template ./auth/buf.gen.yaml --path ./auth/v1/auth.proto


build_auth_grpc_svc:
	cd ./service/gRPC/auth && go build -o ./build/auth_grpc_service ./cmd/server/main.go &&\
		go build -o ./build/auth_gateway_service ./cmd/gateway/main.go
