IMAGE = output-local

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

binary:
	$(MAKE) DOCKERFILE=Dockerfile.after TAG=after EXTRA_PARAMS="--target=binary --output=type=local,dest=$(PWD)"

helm-chart:
	$(MAKE) DOCKERFILE=Dockerfile.after2 TAG=after EXTRA_PARAMS="--target=helm-chart --output=type=local,dest=$(PWD)"

generate:
	$(MAKE) DOCKERFILE=Dockerfile.after3 TAG=after EXTRA_PARAMS="--target=codegen --output=type=local,dest=$(PWD)"

unit-tests:
	$(MAKE) DOCKERFILE=Dockerfile.after4 TAG=after EXTRA_PARAMS="--target=unit-tests --output=type=local,dest=$(PWD) --progress=plain"

prune:
	docker buildx prune -a

clean:
	rm -rf build/
