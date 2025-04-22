user_srv:
	go run ./srvs/user_srv/*.go -port 40000
user_srv_idl:
	cd ./srvs/user_srv/proto && protoc --go_out=./gen --go-grpc_out=./gen user.proto
user_web_idl:
	cd ./apis/user_web && hz update -idl ./proto/user_web.proto   
user_web:
	go run ./apis/user_web/*.go -port 8001

homestay_srv_idl:
	cd ./srvs/homestay_srv/proto && protoc --go_out=./gen --go-grpc_out=./gen homestay.proto
homestay_srv:
	go run ./srvs/homestay_srv/*.go -port 40001

homestay_web_idl:
	cd ./apis/homestay_web && hz update -idl ./proto/homestay_web.proto
homestay_web:
	go run ./apis/homestay_web/*.go -port 8002

order_srv_idl:
	cd ./srvs/order_srv/proto && protoc --go_out=./gen --go-grpc_out=./gen order.proto
order_srv:
	go run ./srvs/order_srv/*.go -port 40002
order_web_idl:
	cd ./apis/order_web && hz update -idl ./proto/order_web.proto
order_web:
	go run ./apis/order_web/*.go -port 8003

payment_srv_idl:
	cd ./srvs/payment_srv/proto && protoc --go_out=./gen --go-grpc_out=./gen payment.proto
payment_srv:
	go run ./srvs/payment_srv/*.go -port 40003

payment_web_idl:
	cd ./apis/payment_web && hz update -idl ./proto/payment_web.proto
payment_web:
	go run ./apis/payment_web/*.go -port 8004

mq:
	go run ./srvs/mq/*.go

asynq:
	go run ./srvs/asynq/*.go
consul:
	consul agent -server -bootstrap-expect 1 -ui -node=consul-dev  -data-dir=/home/yjy/data/consul -config-dir=/home/yjy/data/consul/config -bind=127.0.0.1