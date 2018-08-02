include ../../scripts/go.mk

# Build as a dockerfile
docker:
	$(MAKE) --directory=../.. --file=scripts/docker.mk TARGET=server DOCKERFILE=services/yfuzz-server/Dockerfile

.PHONY: docker
