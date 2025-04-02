image-build:
	docker buildx build -f $(DOCKERFILE) -t mount-cache-zypper:$(TAG) $(EXTRA_PARAMS) .

fast:
	$(MAKE) DOCKERFILE=Dockerfile.fast TAG=fast

slow:
	$(MAKE) DOCKERFILE=Dockerfile.slow TAG=slow

prune:
	docker buildx prune -a
