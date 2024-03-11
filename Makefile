.PHONY: test e2e-test cover gofmt gofmt-fix header-check clean tar.gz docker-push release docker-push-all flannel-git

# Default tag and architecture. Can be overridden
TAG?=$(shell git describe --tags --dirty --always)
ifeq ($(TAG),)
	TAG=latest
endif

#ifeq ($(findstring dirty,$(TAG)), dirty)
#    TAG=latest
#endif

### BUILDING
#debug:
#	@echo $(TAG)
clean:
	rm -f ninja

ninja: $(shell find . -type f  -name '*.go')
	go build -o ninja \
	  -ldflags '-s -w -X "github.com/daddvted/fruitninja/fruitninja.Version=$(TAG)" -extldflags "-static"'