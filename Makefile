MAKE_PATH=$(GOPATH)/bin:/bin:/usr/bin:/usr/local/bin:$PATH
BIN_PATH=/home/dm/App/go/bin
BUF_VERSION=1.50.1

.PHONY: buf-install
buf-install:
ifeq ($(shell uname -s), Darwin)
	@[ ! -f $(BIN_PATH)/buf ] || exit 0 && \
	tmp=$$(mktemp -d) && cd $$tmp && \
	curl -L -o buf \
		https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-Darwin-arm64 && \
	mv buf $(BIN_PATH)/buf && \
	sudo chmod +x $(BIN_PATH)/buf
else
	@[ ! -f $(BIN_PATH)/buf ] || exit 0 && \
	tmp=$$(mktemp -d) && cd $$tmp && \
	curl -L -o buf \
		https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-Linux-x86_64 && \
	mv buf $(BIN_PATH)/buf && \
	sudo chmod +x $(BIN_PATH)/buf
endif

.PHONY: gen
gen:
	buf generate
	@for dir in $(CURDIR)/gen/go/*/; do \
	  cd $$dir && \
	  go mod init && go mod tidy; \
  	done

.PHONY: lint
lint:
	@$(BIN_PATH)/buf lint --config buf.yaml

.PHONY: format
format:
	@$(BIN_PATH)/buf format


.PHONY: clean
clean:
	@rm -rf ./gen || true