VERSION = dev
TAG = plutono/build-container
USER_ID = $(shell id -u)
GROUP_ID = $(shell id -g)

all: build deploy

build:
	docker build -t "${TAG}:${VERSION}" .

deploy:
	docker push "${TAG}:${VERSION}"

run:
	docker run -ti \
		-e "CIRCLE_BRANCH=local" \
		-e "CIRCLE_BUILD_NUM=472" \
		${TAG}:${VERSION} \
		bash

run-with-local-source-live:
	docker run -d \
		-e "CIRCLE_BRANCH=local" \
		-e "CIRCLE_BUILD_NUM=472" \
		-w "/go/src/github.com/credativ/plutono" \
		--name plutono-build \
		-v "${GOPATH}/src/github.com/credativ/plutono:/go/src/github.com/credativ/plutono" \
		${TAG}:${VERSION} \
		bash -c "/tmp/bootstrap.sh; mkdir /.cache; chown "${USER_ID}:${GROUP_ID}" /.cache; tail -f /dev/null"
	docker exec -ti --user "${USER_ID}:${GROUP_ID}" plutono-build bash

run-with-local-source-copy:
	docker run -d \
		-e "CIRCLE_BRANCH=local" \
		-e "CIRCLE_BUILD_NUM=472" \
		-w "/go/src/github.com/credativ/plutono" \
		--name plutono-build \
		${TAG}:${VERSION} \
		bash -c "/tmp/bootstrap.sh; tail -f /dev/null"
	docker cp "${GOPATH}/src/github.com/credativ/plutono" plutono-build:/go/src/github.com/grafana/
	docker exec -ti plutono-build bash

update-source:
	docker cp "${GOPATH}/src/github.com/credativ/plutono" plutono-build:/go/src/github.com/grafana/

attach:
	docker exec -ti plutono-build bash

attach-live:
	docker exec -ti --user "${USER_ID}:${GROUP_ID}" plutono-build bash

stop:
	docker kill plutono-build
	docker rm plutono-build
