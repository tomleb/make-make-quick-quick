image-build:
	docker buildx build -f $(DOCKERFILE) -t mount-cache-go:$(TAG) $(EXTRA_PARAMS) .

fast:
	$(MAKE) DOCKERFILE=Dockerfile.fast TAG=fast

slow:
	$(MAKE) DOCKERFILE=Dockerfile.slow TAG=slow

fastdebug:
	$(MAKE) DOCKERFILE=Dockerfile.fastdebug TAG=fastdebug EXTRA_PARAMS="--progress=plain"

prune:
	docker buildx prune -a
