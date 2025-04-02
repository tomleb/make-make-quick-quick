IMAGE = copy-from

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

fast-2.10:
	$(MAKE) DOCKERFILE=Dockerfile.fast TAG=fast EXTRA_PARAMS="--build-arg=CHART_DEFAULT_BRANCH=dev-v2.10"

fast-2.9:
	$(MAKE) DOCKERFILE=Dockerfile.fast TAG=fast EXTRA_PARAMS="--build-arg=CHART_DEFAULT_BRANCH=dev-v2.9"

slow:
	$(MAKE) DOCKERFILE=Dockerfile.slow TAG=slow

slow-2.10:
	$(MAKE) DOCKERFILE=Dockerfile.slow TAG=slow EXTRA_PARAMS="--build-arg=CHART_DEFAULT_BRANCH=dev-v2.10"

slow-2.9:
	$(MAKE) DOCKERFILE=Dockerfile.slow TAG=slow EXTRA_PARAMS="--build-arg=CHART_DEFAULT_BRANCH=dev-v2.9"

prune:
	docker buildx prune -a
