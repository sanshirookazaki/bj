COVERALLS_TOKEN=

.PHONY: test
test:
	go test -v ./...

.PHONY: build
build:
	go build -v

ci:
	circleci config validate
	circleci local execute -e COVERALLS_TOKEN=$(COVERALLS_TOKEN)