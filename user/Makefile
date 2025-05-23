include vendor.proto.mk

# Используем bin в текущей директории для установки плагинов protoc
LOCAL_BIN := $(CURDIR)/bin

# Добавляем bin в текущей директории в PATH при запуске protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

# Путь до protobuf файлов
PROTO_PATH := $(CURDIR)/api

# Путь до сгенеренных .pb.go файлов
PKG_PROTO_PATH := $(CURDIR)/pkg

# Путь до завендореных protobuf файлов
VENDOR_PROTO_PATH := $(CURDIR)/vendor.protobuf

# устанавливаем необходимые плагины
.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	$(info Installing binary dependencies...)

	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/bufbuild/buf/cmd/buf@v1.32.2
	go install github.com/yoheimuta/protolint/cmd/protolint@latest

generate: export GOBIN := $(LOCAL_BIN)
generate:
	PATH=$(LOCAL_BIN):PATH $(LOCAL_BIN)/buf generate
    
# Билд приложения
build:
	go build -o $(LOCAL_BIN) ./cmd/notes/client
	go build -o $(LOCAL_BIN) ./cmd/notes/server

# Линтер
lint: .proto-lint

# Линтер proto файлов
.proto-lint:
	$(LOCAL_BIN)/protolint -config_path ./.protolint.yaml ./proto/

# Форматирование protobuf файлов
proto-format:
	$(info run buf format...)
	$(LOCAL_BIN)/buf format -w
	
# Объявляем, что текущие команды не являются файлами и
# интсрументируем Makefile не искать изменения в файловой системе
.PHONY: \
	.bin-deps \
	.protoc-generate \
	.tidy \
	.proto-lint \
	proto-format \
	vendor \
	generate \
	build \
	lint
