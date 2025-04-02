IMAGE = mount-cache-go

image-build:
	@echo "####################################################"
	@echo "#"
	@echo "# Running the following command:"
	@echo "#   docker buildx build -f $(DOCKERFILE) -t $(IMAGE):$(TAG) $(EXTRA_PARAMS) ."
	@echo "#"
	@echo "####################################################"
	@echo ""
	@docker buildx build -f $(DOCKERFILE) -t $(IMAGE):$(TAG) $(EXTRA_PARAMS) .

fast:
	$(MAKE) DOCKERFILE=Dockerfile.fast TAG=fast

slow:
	$(MAKE) DOCKERFILE=Dockerfile.slow TAG=slow

fastdebug:
	$(MAKE) DOCKERFILE=Dockerfile.fastdebug TAG=fastdebug EXTRA_PARAMS="--progress=plain"

prune:
	docker buildx prune -a
