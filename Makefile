# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
PROTOCMD=protoc
DOCKER_CMD=docker
DOCKER_FILE="./Dockerfile"
DOCKER_PATH="./"


gen-proto:
	$(PROTOCMD) -I=. -I=$(GOPATH)/src --go_out=plugins=grpc:. ./app/responder/service.proto

build-linux:
		CGO_ENABLED=0 \
		GOOS=linux \
			$(GOBUILD) \
				-installsuffix cgo \
				-o ./linux/bin/server \
				-ldflags "-X main.Version=$(APP_VERSION)" \
				./bin/server/main.go;
		CGO_ENABLED=0 \
        GOOS=linux \
        	$(GOBUILD) \
        		-installsuffix cgo \
        		-o ./linux/bin/client \
        		-ldflags "-X main.Version=$(APP_VERSION)" \
        		./bin/client/main.go;

build-docker: build-linux
		$(DOCKER_CMD) build -t dvcdsys/examples:grpcbalancing -f $(DOCKER_FILE) $(DOCKER_PATH)
		docker push dvcdsys/examples:grpcbalancing
