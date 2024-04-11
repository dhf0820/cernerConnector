USER ?= dhf0820
NS ?= dhf0820
TAG ?= latest
TEST = dhf0820
PROD = vertisoft
VERSION ?= $(TAG)
ARC ?= amd64
ARCH ?= $(ARC)
IMG_NAME ?= cerner_connector
ARM_IMG_NAME ?= cerner_conn_arm 
AMD_IMG_NAME ?= cerner_connector 
IMAGE_NAME ?= $(IMG_NAME)
IMAGE_TEST_NAME ?= $(IMAGE_NAME)_test
LINUX_IMAGE_NAME ?= $(IMAGE_NAME)
LINUX_TEST_IMAGE_NAME ?= $(IMAGE_NAME)_test
BINARY_NAME=$(IMAGE_NAME)
BINARY_UNIX=$(BINARY_NAME)_linux
BINARY_TEST_UNIX=$(BINARY_NAME)_test
M1_IMAGE_NAME ?= $(IMG_NAME)_$(ARCH)
M1_TEST_IMAGE_NAME ?= $(M1_IMAGE_NAME)_test
M1_UNIX=$(BINARY_NAME)
BINARY_TEST_M1=$(BINARY_NAME)_m1
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test -run
GOGET=$(GOCMD) get
BINARY=$(IMG_NAME)_$(ARCH)
DOCKER_NAME=$(IMG_NAME)_$(ARCH)
MAC_IMAGE_NAME= $(BINARY)


.PHONY: all api server client cert pb1

all: server client


#api/api.pb.go:
#protoc -I ./ --proto_path=./ --go_out=./ pkg/proto/*.proto
api:
	@protoc -I ./protobufs/ \
		--proto_path=./ \
		--go_out=plugins=grpc:./ \
		./protobufs/*.proto


#	@protoc -I ./protobufs \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--proto_path=./ \
		--go_out=plugins=grpc:./ \
		./protobufs/*.proto



#api: api/api.pb.go api/api.pb.gw.go api/api.swagger.json ## Auto-generate grpc go sources


dep: ## Get the dependencies
	@go get -v -d ./...

tidy: # add all new includes
	@go mod tidy

#build:
#	$(GOBUILD) -o $(BINARY_NAME) -v

build_mac:
	ARCH=arm64
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) -o $(BINARY_NAME) -v
	#docker build -t $(TEST)/$(DOCKER_NAME):$(VERSION) -f Dockerfile_$(ARCH) .
	#CGO_ENABLED=0 $(GOBUILD) -o $(BINARY_MAC) -v

#build:
#	ARCH=arm64	
#	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) -o $(BINARY) -v
#	docker build -t $(TEST)/$(DOCKER_NAME):$(VERSION) -f Dockerfile_$(ARCH) .
#	#docker push $(TEST)/$(DOCKER_NAME):$(VERSION)

# push:
#         @docker image tag $(image_name):$(TAG) $(NS)/$(image_name):$(TAG)
#         @docker image push $(NS)/$(image_name):$(TAG)

build:
	ARCH=amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) -o $(BINARY_NAME) 
#  go build -o $(IMAGE_NAME)_amd64
	docker build -t $(TEST)/$(IMAGE_NAME):$(VERSION) -f Dockerfile .
	docker push $(NS)/$(IMAGE_NAME):$(VERSION)


local_test:
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) -o $(IMAGE_NAME) 
#  go build -o $(IMAGE_NAME)_amd64
	docker build -t $(TEST)/$(IMAGE_NAME):$(VERSION) -f Dockerfile_$(ARCH) .
	#docker push $(NS)/$(IMAGE_NAME):$(VERSION)

test:
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) -o $(IMAGE_NAME) 
	docker build -t $(TEST)/$(DOCKER_NAME):$(VERSION) -f Dockerfile_$(ARCH) .
	docker push $(TEST)/$(DOCKER_NAME):$(VERSION)

test_amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(AMD_IMG_NAME)
	docker build -t $(TEST)/$(AMD_IMG_NAME):$(VERSION) -f Dockerfile .
	docker push $(TEST)/$(AMD_IMG_NAME):$(VERSION)

# test:
# 	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
# 	docker build -t $(TEST)/$(IMAGE_NAME):$(VERSION) -f Dockerfile .
# 	docker push $(TEST)/$(IMAGE_NAME):$(VERSION)

prod:
	#CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) -o $(BINARY) -v
	#docker build -t $(PROD)/$(DOCKER_NAME):$(VERSION) -f Dockerfile_amd64 .
	#docker push $(PROD)/$(DOCKER_NAME):$(VERSION)

	 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
	 docker build -t $(PROD)/$(IMAGE_NAME):$(VERSION) -f Dockerfile .
	 docker push $(PROD)/$(IMAGE_NAME):$(VERSION)

release:
	CGO_ENABLED=0 GOOS=linux amd64 $(GOBUILD) -o $(BINARY) -v
	docker build -t $(PROD)/$(DOCKER_NAME):$(VERSION) -f Dockerfile .
	docker push $(PROD)/$(DOCKER_NAME):$(VERSION)

	# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
	# docker build -t $(NS)/$(IMAGE_NAME):$(VERSION) -f Dockerfile .
	# docker push $(NS)/$(IMAGE_NAME):$(VERSION)

core:
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) -o $(BINARY) -v
	docker build -t $(NS)/$(DOCKER_NAME):$(VERSION) -f Dockerfile .
	docker push $(NS)/$(DOCKER_NAME):$(VERSION)

	# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
	# docker build -t $(NS)/$(IMAGE_NAME):$(VERSION) -f Dockerfile .
	# //docker push $(NS)/$(IMAGE_NAME):$(VERSION)

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

docker-build:
	docker build -t $(NS)/$(LINUX_NAME):$(VERSION) -f Dockerfile .

docker-push: # push to docker
	docker push $(NS)/$(IMAGE_NAME):$(VERSION)

client: dep api ## Build the binary file for client
	@go build -i -v -o $(CLIENT_OUT) $(CLIENT_PKG_BUILD)

clean: ## Remove previous builds
	@rm $(SERVER_OUT) $(CLIENT_OUT) $(PB_OUT) $(API_REST_OUT) $(API_SWAG_OUT)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run-server:
	go run main.go -port 8080

run-server-tls:
	go run cmd/server/main.go -port 9901 -tls server

run-server-mutual-tls:
	go run cmd/server/main.go -port 7777 -tls mutual

run-client-do3:
	go run src/client/main.go  -address docker-test.vertisoft.com -port 8080

run-client:
	go run src/client/main.go  -address localhost -port 9001

run-client-do-test:
	go run src/client/main.go  -address 161.35.229.137 -port 30001

run-client-tls:
	go run cmd/client/main.go  -address 0.0.0.0:7777 -tls server

run-client-mutual-tls:
	go run cmd/client/main.go  -address 0.0.0.0:7777 -tls mutual

cert:
	cd cert; ./gen.sh; cd ..

