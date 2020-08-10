SHELL 	   := $(shell which bash)

## BOF define block

BINARIES   := wc
BINARY     = $(word 1, $@)

PLATFORMS  := windows linux darwin
PLATFORM   = $(word 1, $@)

ROOT_DIR   := $(shell git rev-parse --show-toplevel)
BIN_DIR    := $(ROOT_DIR)/bin
REL_DIR    := $(ROOT_DIR)/release
SRC_DIR    := $(ROOT_DIR)/cmd/wc
PPROF_DIR  := $(ROOT_DIR)/pprof
BENCH_DIR  := $(ROOT_DIR)/bench

VERSION    :=`git describe --tags 2>/dev/null`
COMMIT     :=`git rev-parse --short HEAD 2>/dev/null`
DATE       :=`date "+%FT%T%z"`

LDBASE     := github.com/gertd/wild/pkg/version
LDFLAGS    := -ldflags "-w -s -X $(LDBASE).ver=${VERSION} -X $(LDBASE).date=${DATE} -X $(LDBASE).commit=${COMMIT}"

GOARCH     ?= amd64
GOOS       ?= $(shell go env GOOS)
GOPROXY    ?= $(shell go env GOPROXY)

LINTER     := $(BIN_DIR)/golangci-lint
LINTVERSION:= v1.28.1

TESTRUNNER := $(BIN_DIR)/gotestsum
TESTVERSION:= v0.4.2

NO_COLOR   :=\033[0m
OK_COLOR   :=\033[32;01m
ERR_COLOR  :=\033[31;01m
WARN_COLOR :=\033[36;01m
ATTN_COLOR :=\033[33;01m

## EOF define block

.PHONY: all
all: deps gen build test

.PHONY: deps
deps:
	@echo -e "$(ATTN_COLOR)==> $@ $(NO_COLOR)"
	@GO111MODULE=on go mod download
	@GOBIN=$(BIN_DIR) go get golang.org/x/perf/cmd/benchstat

.PHONY: gen
gen:
	@echo -e "$(ATTN_COLOR)==> $@ $(NO_COLOR)"
	@go generate ./...

.PHONY: dobuild
dobuild: deps
	@echo -e "$(ATTN_COLOR)==> $@ $(B) GOOS=$(P) GOARCH=$(GOARCH) VERSION=$(VERSION) COMMIT=$(COMMIT) DATE=$(DATE) $(NO_COLOR)"
	@GOOS=$(P) GOARCH=$(GOARCH) GO111MODULE=on go build $(LDFLAGS) -o $(T)/$(P)-$(GOARCH)/$(B)$(if $(findstring $(P),windows),".exe","") $(SRC_DIR)
ifneq ($(P),windows)
	@chmod +x $(T)/$(P)-$(GOARCH)/$(B)
endif

.PHONY: build 
build: $(BIN_DIR)
	@echo -e "$(ATTN_COLOR)==> $@ $(NO_COLOR)"
	@for b in ${BINARIES}; 									\
	do 														\
		$(MAKE) dobuild B=$${b} P=${GOOS} T=${BIN_DIR}; 	\
	done 													

.PHONY: doinstall
doinstall:
	@echo -e "$(ATTN_COLOR)==> $@ GOOS=$(P) GOARCH=$(GOARCH) VERSION=$(VERSION) COMMIT=$(COMMIT) DATE=$(DATE) $(NO_COLOR)"
	@GOOS=$(P) GOARCH=$(GOARCH) GO111MODULE=on go install $(LDFLAGS) $(SRC_DIR)

.PHONY: install
install: 
	@echo -e "$(ATTN_COLOR)==> $@ $(NO_COLOR)"
	@for b in ${BINARIES}; 									\
	do 														\
		$(MAKE) doinstall B=$${b} P=${GOOS}; 			 	\
	done 													

.PHONY: dorelease
dorelease:
	@echo -e "$(ATTN_COLOR)==> $@ $(B) GOOS=$(P) GOARCH=$(GOARCH) VERSION=$(VERSION) COMMIT=$(COMMIT) DATE=$(DATE) $(NO_COLOR)"
	@GOOS=$(P) GOARCH=$(GOARCH) GO111MODULE=on go build $(LDFLAGS) -o $(T)/$(P)-$(GOARCH)/$(B)$(if $(findstring $(P),windows),".exe","") $(SRC_DIR)
ifneq ($(P),windows)
	@chmod +x $(T)/$(P)-$(GOARCH)/$(B)
endif

.PHONY: release
release: $(REL_DIR)
	@echo -e "$(ATTN_COLOR)==> $@ $(NO_COLOR)"
	@for b in ${BINARIES}; 									\
	do 														\
		for p in ${PLATFORMS};								\
		do 													\
			$(MAKE) dorelease B=$${b} P=$${p} T=${REL_DIR};	\
		done;												\
	done 													\

$(REL_DIR):
	@echo -e "$(ATTN_COLOR)==> create REL_DIR $(REL_DIR) $(NO_COLOR)"
	@mkdir -p $(REL_DIR)

$(BIN_DIR):
	@echo -e "$(ATTN_COLOR)==> create BIN_DIR $(BIN_DIR) $(NO_COLOR)"
	@mkdir -p $(BIN_DIR)

$(PPROF_DIR):
	@echo -e "$(ATTN_COLOR)==> create PPROF_DIR $(PPROF_DIR) $(NO_COLOR)"
	@mkdir -p $(PPROF_DIR)

$(BENCH_DIR):
	@echo -e "$(ATTN_COLOR)==> create PPROF_DIR $(BENCH_DIR) $(NO_COLOR)"
	@mkdir -p $(BENCH_DIR)

$(TESTRUNNER):
	@echo -e "$(ATTN_COLOR)==> get $@  $(NO_COLOR)"
	@GOBIN=$(BIN_DIR) go get -u gotest.tools/gotestsum 

.PHONY: test 
test: $(TESTRUNNER)
	@echo -e "$(ATTN_COLOR)==> $@ $(NO_COLOR)"
	@CGO_ENABLED=0 $(BIN_DIR)/gotestsum --format short-verbose -- -count=1 -v $(ROOT_DIR)/...

$(LINTER):
	@echo -e "$(ATTN_COLOR)==> get $@  $(NO_COLOR)"
	@curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s $(LINTVERSION)
 
.PHONY: lint
lint: $(LINTER)
	@echo -e "$(ATTN_COLOR)==> $@ $(NO_COLOR)"
	@CGO_ENABLED=0 $(LINTER) run --enable-all --disable gofumpt,noctx
	@echo -e "$(NO_COLOR)\c"

.PHONY: clean
clean:
	@echo -e "$(ATTN_COLOR)==> $@ $(NO_COLOR)"
	@rm -rf $(BIN_DIR)
	@rm -rf $(REL_DIR)
	@rm -rf $(PPROF_DIR)
	@rm -rf $(BENCH_DIR)

.PHONY: bench
bench: $(BENCH_DIR)
	@echo -e "$(ATTN_COLOR)==> $@ $(NO_COLOR)"
	
	@echo -e "$(ATTN_COLOR) running BenchmarkFilepathMatch (BASELINE) $(NO_COLOR)"
	@go test -count 5 -run=NONE  -benchmem  -bench=BenchmarkFilepathMatch . | sed 's/BenchmarkFilepathMatch*/BenchmarkMatch/g' > $(BENCH_DIR)/wild-fpmatch.txt
	
	@echo -e "$(ATTN_COLOR) running BenchmarkMatch $(NO_COLOR)"
	@go test -count 5 -run=NONE  -benchmem  -bench=BenchmarkMatch . > $(BENCH_DIR)/wild-match.txt
	
	@echo -e "$(ATTN_COLOR) create report $(BENCH_DIR)/report.txt $(NO_COLOR)"
	@$(BIN_DIR)/benchstat $(BENCH_DIR)/wild-fpmatch.txt $(BENCH_DIR)/wild-match.txt > $(BENCH_DIR)/report.txt

.PHONY: profile
profile: $(PPROF_DIR)
	@echo -e "$(ATTN_COLOR)==> $@ $(NO_COLOR)"
	@go test -run='^\$$' -bench=. -cpuprofile=$(PPROF_DIR)/cpu.pprof -benchmem -memprofile=$(PPROF_DIR)/mem.pprof -trace $(PPROF_DIR)/trace.out ./cmd
	@go tool trace -pprof=net $(PPROF_DIR)/trace.out > $(PPROF_DIR)/net.pprof
	@go tool trace -pprof=sync $(PPROF_DIR)/trace.out > $(PPROF_DIR)/sync.pprof
	@go tool trace -pprof=syscall $(PPROF_DIR)/trace.out > $(PPROF_DIR)/syscall.pprof
	@go tool trace -pprof=sched $(PPROF_DIR)/trace.out > $(PPROF_DIR)/sched.pprof

.PHONY: memprofile
memprofile:
	@echo -e "$(ATTN_COLOR)==> $@ $(NO_COLOR)"
	@go tool pprof -http=:8080 ./pprof/mem.pprof

.PHONY: cpuprofile
cpuprofile:
	@echo -e "$(ATTN_COLOR)==> $@ $(NO_COLOR)"
	@go tool pprof -http=:8080 ./pprof/cpu.pprof

