include vendor.proto.mk

# Используем bin в текущей директории для установки плагинов protoc
LOCAL_BIN := $(CURDIR)/bin

# Добавляем bin в текущей директории в PATH при запуске protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

# Путь до protobuf файлов
PROTO_PATH := $(CURDIR)/proto

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

# генерация .go файлов с помощью protoc
.protoc-generate:
	$(info run protoc...)

	mkdir -p $(PKG_PROTO_PATH)
	$(PROTOC) -I $(VENDOR_PROTO_PATH) --proto_path=$(CURDIR) \
	--go_out=$(PKG_PROTO_PATH) --go_opt paths=source_relative \
	--go-grpc_out=$(PKG_PROTO_PATH) --go-grpc_opt paths=source_relative \
	--grpc-gateway_out=$(PKG_PROTO_PATH) --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true \
	$(PROTO_PATH)/messages.proto $(PROTO_PATH)/service.proto
	
	$(PROTOC) -I $(VENDOR_PROTO_PATH) --proto_path=$(CURDIR) \
	--openapiv2_out=. --openapiv2_opt logtostderr=true \
	$(PROTO_PATH)/service.proto

# Генерация .pb файлов с помощью buf
.buf-generate:
	$(info run buf generate...)
	PATH="$(LOCAL_BIN):$(PATH)" $(LOCAL_BIN)/buf generate

# Форматирование protobuf файлов
proto-format:
	$(info run buf format...)
	$(LOCAL_BIN)/buf format -w

# go mod tidy
.tidy:
	GOBIN=$(LOCAL_BIN) go mod tidy

# Генерация кода из protobuf
generate: .bin-deps .buf-generate proto-format .tidy


# Сборка приложения
build:
CGO_ENABLED=0 go build -v -o $(LOCAL_BIN) ./cmd/server/main.go

# Линтер
lint: .proto-lint

# Линтер proto файлов
.proto-lint:
	$(LOCAL_BIN)/protolint -config_path ./.protolint.yaml ./proto/

# docker-compose up
up:
	docker-compose up --build -d

# docker-compose down
down:
	docker-compose down

# Объявляем, что текущие команды не являются файлами и
# интсрументируем Makefile не искать изменения в файловой системе
.PHONY: \
	.bin-deps \
	.protoc-generate \
	.buf-generate \
	.tidy \
	.proto-lint \
	proto-format \
	generate \
	build \
	lint \
	up \
	down