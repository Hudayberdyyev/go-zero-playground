# API Gateway
gen-api:
	goctl api -o api/gateway.api
gen-api-code:
	goctl api go -api api/gateway.api -dir api/.
run-gateway:
	go run api/gateway.go -f api/etc/gateway-api.yaml

# Trending
#gen-trending:
#	goctl rpc -o rpc/trending/trending.proto
#gen-trending-code:
#	goctl rpc protoc rpc/trending/trending.proto --go_out=./rpc/trending --go-grpc_out=./rpc/trending --zrpc_out=./rpc/trending
#run-trending:
#	go run rpc/trending/trending.go -f rpc/trending/etc/trending.yaml
