gen:
	protoc --go_out=infra/pb --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=infra/pb proto/*.proto 
clean:
	rm -r infra/pb/*.go