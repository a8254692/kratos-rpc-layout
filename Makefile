#GOPATH:=$(shell go env GOPATH)
#VERSION=$(shell git describe --tags --always)

PROTO_FILES=$(shell find ./api -name *.proto)


.PHONY: init
# init ENV
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/google/wire/cmd/wire@latest
	go install entgo.io/ent/cmd/ent@latest


.PHONY: generate
# generate client code
generate:
	go generate ./...

.PHONY: sql2file
sql2file:
	mysqldump --compact --skip-add-drop-table -d -h 192.168.152.129 -P 3306 -u root -p user > ./api/sql/user.sql;

.PHONY: sql2ent
sql2ent:
	sql2ent mysql ddl -src "./api/sql/*.sql" -dir "./internal/data/ent/schema"

.PHONY: proto
# generate internal proto
proto:
	protoc	-I ./third_party  \
 			--proto_path=. \
        	--go_out=paths=source_relative:. \
        	--go-grpc_opt=require_unimplemented_servers=false \
        	--go-grpc_out=paths=source_relative:. \
        	$(PROTO_FILES)

.PHONY: ent
ent:
	ent generate ./internal/data/ent/schema



#.PHONY: grpc
## generate grpc code
#grpc:
#	protoc --proto_path=. \
#           --proto_path=$(KRATOS)/third_party \
#           --go_out=paths=source_relative:. \
#           --go-grpc_out=paths=source_relative:. \
#           $(API_PROTO_FILES)
#
#.PHONY: http
## generate http code
#http:
#	protoc --proto_path=. \
#           --proto_path=$(KRATOS)/third_party \
#           --go_out=paths=source_relative:. \
#           --go-http_out=paths=source_relative:. \
#           $(API_PROTO_FILES)
#
#.PHONY: errors
## generate errors code
#errors:
#	protoc --proto_path=. \
#           --proto_path=$(KRATOS)/third_party \
#           --go_out=paths=source_relative:. \
#           --go-errors_out=paths=source_relative:. \
#           $(API_PROTO_FILES)
#
#.PHONY: proto
## generate internal proto
#proto:
#	protoc --proto_path=. \
#           --proto_path=$(KRATOS)/third_party \
#           --go_out=paths=source_relative:. \
#           $(INTERNAL_PROTO_FILES)
#
#.PHONY: swagger
## generate swagger file
#swagger:
#	protoc --proto_path=. \
#		--proto_path=$(KRATOS)/third_party \
#		--openapiv2_out . \
#		--openapiv2_opt logtostderr=true \
#		$(API_PROTO_FILES)
#
#.PHONY: build
## build
#build:
#	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...
#
#.PHONY: test
## test
#test:
#	go test -v ./... -cover
#
#.PHONY: all
## generate all
#all:
#	make generate;
#	make grpc;
#	make http;
#	make proto;
#	make errors;
#	make swagger;
#	make build;
#	make test;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
