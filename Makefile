
docker/release: docker/base
	# push the base image
	docker push cohix/goott-server-base:$(shell goott-server version)
	docker push cohix/goott-server-base:latest
	# build and push the server
	docker build --pull --no-cache . -t cohix/goott-server:$(shell goott-server version)
	docker push cohix/goott-server:$(shell goott-server version)

docker/dev:
	docker build . -t cohix/goott-server:dev
	docker run -it cohix/goott-server:dev

docker/base:
	go install
	docker build ./ops/base-image -t "cohix/goott-server-base:$(shell goott-server version)"
	docker tag "cohix/goott-server-base:$(shell goott-server version)" cohix/goott-server-base:latest

proto/model:
	protoc -I=../ -I=. -I=model/proto --go_out=plugins=grpc:$(GOPATH)/src $(shell ls ./model/proto/)

proto/service:
	protoc -I=../ -I=. -I=service/proto --go_out=plugins=grpc:$(GOPATH)/src $(shell ls ./service/proto/)

proto: proto/model proto/service

.PHONY: proto
	