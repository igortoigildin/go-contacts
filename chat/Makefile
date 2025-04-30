.PHONY: proto-gen clean

# Используем bin в текущей директории для установки плагинов protoc
LOCAL_BIN := $(CURDIR)/bin

# Добавляем bin в текущей директории в PATH при запуске protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

# Путь до protobuf файлов
PROTO_PATH := $(CURDIR)/api/proto

# Путь до сгенеренных .pb.go файлов
PKG_PROTO_PATH := $(CURDIR)/pkg

# устанавливаем необходимые плагины
bin-deps: export GOBIN := $(LOCAL_BIN)
bin-deps:
	$(info Installing binary dependencies...)
	
	# Create bin directory if it doesn't exist
	mkdir -p $(LOCAL_BIN)
	
	# Install protoc for macOS
	brew install protobuf
	
	# Install protoc plugins with explicit paths
	go install -v google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install -v github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	
	# Make plugins executable
	chmod +x $(LOCAL_BIN)/protoc-gen-go
	chmod +x $(LOCAL_BIN)/protoc-gen-go-grpc
	chmod +x $(LOCAL_BIN)/protoc-gen-openapiv2
	
	# Verify installation
	@echo "Verifying installation..."
	@which protoc || (echo "protoc not found" && exit 1)
	@ls -l $(LOCAL_BIN)/protoc-gen-go || (echo "protoc-gen-go not found" && exit 1)
	@ls -l $(LOCAL_BIN)/protoc-gen-go-grpc || (echo "protoc-gen-go-grpc not found" && exit 1)
	@ls -l $(LOCAL_BIN)/protoc-gen-openapiv2 || (echo "protoc-gen-openapiv2 not found" && exit 1)

# генерация .go файлов с помощью protoc
protoc-generate:
	mkdir -p $(PKG_PROTO_PATH)
	$(PROTOC) --plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
		--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
		--plugin=protoc-gen-openapiv2=$(LOCAL_BIN)/protoc-gen-openapiv2 \
		--proto_path=$(CURDIR) \
		--proto_path=$(GOPATH)/src/github.com/googleapis/googleapis \
		--go_out=$(PKG_PROTO_PATH) --go_opt paths=source_relative \
		--go-grpc_out=$(PKG_PROTO_PATH) --go-grpc_opt paths=source_relative \
		--openapiv2_out=$(PKG_PROTO_PATH) --openapiv2_opt logtostderr=true \
		$(PROTO_PATH)/messages.proto \
		$(PROTO_PATH)/service.proto

# go mod tidy
tidy:
	GOBIN=$(LOCAL_BIN) go mod tidy

# Генерация кода из protobuf
generate: bin-deps protoc-generate tidy

# Билд приложения
build:
	go build -o $(LOCAL_BIN) ./cmd/main.go
	go build -o $(LOCAL_BIN) ./cmd/client.go
