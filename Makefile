export JAEGER_ENDPOINT=http://localhost:14268/api/traces
export JAEGER_AGENT_HOST=localhost

.PHONY: build
build:
	go build -o clarityai security_api.go

.PHONY: run
run:
	./clarityai

.PHONY: jaeger
jaeger:
	docker run --rm -p 16686:16686 jaegertracing/all-in-one:1.28

.PHONY: test
test:
	curl http://localhost:8080/

.PHONY: clean
clean:
	rm -f clarityai

