up_roosh:
	docker run -it --rm -v ${PWD}\:/go/src/roosh-app -v ${GOPATH}/pkg/mod\:/go/pkg/mod -p 8080\:8080  --name roosh-app roosh-app ./run.sh

build_roosh:
	docker build -t roosh-app .
