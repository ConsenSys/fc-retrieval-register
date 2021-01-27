REGISTRY?=consensys/
VERSION?=dev

dev:
	docker-compose -f docker-compose.dev.yml up

# stop:
# 	docker-compose stop

build:
	docker build -f Dockerfile.dev -t ${REGISTRY}fc-retrieval-register:${VERSION} .

build-dev:
	go build -v cmd/register-server/main.go