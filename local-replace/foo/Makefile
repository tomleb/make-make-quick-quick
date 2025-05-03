IMAGE = local-replace

PLAIN ?= false

BAR_DEP_PATH = $(shell go list -json -m all | jq 'select(.Path == "github.com/tomleb/make-make-quick-quick/local-replace/bar" and has("Replace")) | .Replace.Dir' -r)

ifneq ($(BAR_DEP_PATH),)
	EXTRA_PARAMS += --build-context=bar-context=$(BAR_DEP_PATH)
	EXTRA_PARAMS += --build-arg=BAR_DEP=bar-context
	EXTRA_PARAMS += --build-arg=BAR_DEP_PATH=$(BAR_DEP_PATH)
	EXTRA_PARAMS += --build-arg=BUILD_WORKDIR=$(PWD)
endif

ifeq ($(PLAIN), true)
	EXTRA_PARAMS += --progress=plain
endif

image-build:
	@echo "####################################################"
	@echo "#"
	@echo "# Running the following command:"
	@echo "#   docker buildx build -f $(DOCKERFILE) -t $(IMAGE):$(TAG) $(EXTRA_PARAMS) ."
	@echo "#"
	@echo "####################################################"
	@echo ""
	@docker buildx build -f $(DOCKERFILE) -t $(IMAGE):$(TAG) $(EXTRA_PARAMS) .

before:
	$(MAKE) DOCKERFILE=Dockerfile.before TAG=before

after:
	$(MAKE) DOCKERFILE=Dockerfile.after TAG=after

prune:
	docker buildx prune -a
