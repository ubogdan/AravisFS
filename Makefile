## Param in dlv debbuger
# -> make debug_adret -- decryptls -key="toto" dezdzede
params = 
ifeq (debug_adret,$(firstword $(MAKECMDGOALS)))
    params = yes
endif
ifeq (debug_ubac,$(firstword $(MAKECMDGOALS)))
    params = yes
endif

ifdef params
  # use the rest as arguments for debug
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

before.build:
	go mod download && go mod vendor

build.adret:
	@echo "build in ${PWD}";go build -o adret cmd/adret/main.go

build.adretctl:
	@echo "build in ${PWD}";go build -o adretctl cmd/adretctl/main.go

build.ubac:
	@echo "build in ${PWD}";go build -o ubac cmd/ubac/main.go

debug.adret:
	@dlv debug github.com/ariary/AravisFS/cmd/adret -- $(RUN_ARGS)

debug.ubac:
	@dlv debug github.com/ariary/AravisFS/cmd/ubac -- $(RUN_ARGS)

test.adret:
	@echo "test adret";go test pkg/test/adret/adret_test.go

test.ubac:
	@echo "test ubac";go test pkg/test/ubac/ubac_test.go
