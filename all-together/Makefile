IMAGE = all-together

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
	$(MAKE) DOCKERFILE=Dockerfile TAG=server EXTRA_PARAMS="--target=server"

server-2.10:
	$(MAKE) DOCKERFILE=Dockerfile TAG=server EXTRA_PARAMS="--target=server --build-arg=CHART_DEFAULT_BRANCH=dev-v2.10"

server-2.9:
	$(MAKE) DOCKERFILE=Dockerfile TAG=server EXTRA_PARAMS="--target=server --build-arg=CHART_DEFAULT_BRANCH=dev-v2.9"

agent:
	$(MAKE) DOCKERFILE=Dockerfile TAG=agent EXTRA_PARAMS="--target=agent"

agent-2.10:
	$(MAKE) DOCKERFILE=Dockerfile TAG=agent EXTRA_PARAMS="--target=agent --build-arg=CHART_DEFAULT_BRANCH=dev-v2.10"

agent-2.9:
	$(MAKE) DOCKERFILE=Dockerfile TAG=agent EXTRA_PARAMS="--target=agent --build-arg=CHART_DEFAULT_BRANCH=dev-v2.9"

prune:
	docker buildx prune -a
