IMAGE = multi-target

image-build:
	@echo "####################################################"
	@echo "#"
	@echo "# Running the following command:"
	@echo "#   docker buildx build -f $(DOCKERFILE) -t $(IMAGE):$(TAG) $(EXTRA_PARAMS) ."
	@echo "#"
	@echo "####################################################"
	@echo ""
	@docker buildx build -f $(DOCKERFILE) -t $(IMAGE):$(TAG) $(EXTRA_PARAMS) .


server:
	$(MAKE) DOCKERFILE=Dockerfile.server TAG=server

server-2.10:
	$(MAKE) DOCKERFILE=Dockerfile.server TAG=server EXTRA_PARAMS="--build-arg=CHART_DEFAULT_BRANCH=dev-v2.10"

server-2.9:
	$(MAKE) DOCKERFILE=Dockerfile.server TAG=server EXTRA_PARAMS="--build-arg=CHART_DEFAULT_BRANCH=dev-v2.9"

agent:
	$(MAKE) DOCKERFILE=Dockerfile.server TAG=server
	$(MAKE) DOCKERFILE=Dockerfile.agent TAG=agent

agent-2.10:
	$(MAKE) DOCKERFILE=Dockerfile.server TAG=server EXTRA_PARAMS="--build-arg=CHART_DEFAULT_BRANCH=dev-v2.10"
	$(MAKE) DOCKERFILE=Dockerfile.agent TAG=agent EXTRA_PARAMS="--build-arg=CHART_DEFAULT_BRANCH=dev-v2.10"

agent-2.9:
	$(MAKE) DOCKERFILE=Dockerfile.server TAG=server EXTRA_PARAMS="--build-arg=CHART_DEFAULT_BRANCH=dev-v2.9"
	$(MAKE) DOCKERFILE=Dockerfile.agent TAG=agent EXTRA_PARAMS="--build-arg=CHART_DEFAULT_BRANCH=dev-v2.9"

multi-server:
	$(MAKE) DOCKERFILE=Dockerfile.multi TAG=server EXTRA_PARAMS="--target=server"

multi-server-2.10:
	$(MAKE) DOCKERFILE=Dockerfile.multi TAG=server EXTRA_PARAMS="--target=server --build-arg=CHART_DEFAULT_BRANCH=dev-v2.10"

multi-server-2.9:
	$(MAKE) DOCKERFILE=Dockerfile.multi TAG=server EXTRA_PARAMS="--target=server --build-arg=CHART_DEFAULT_BRANCH=dev-v2.9"

multi-agent:
	$(MAKE) DOCKERFILE=Dockerfile.multi TAG=agent EXTRA_PARAMS="--target=agent"

multi-agent-2.10:
	$(MAKE) DOCKERFILE=Dockerfile.multi TAG=agent EXTRA_PARAMS="--target=agent --build-arg=CHART_DEFAULT_BRANCH=dev-v2.10"

multi-agent-2.9:
	$(MAKE) DOCKERFILE=Dockerfile.multi TAG=agent EXTRA_PARAMS="--target=agent --build-arg=CHART_DEFAULT_BRANCH=dev-v2.9"

prune:
	docker buildx prune -a
