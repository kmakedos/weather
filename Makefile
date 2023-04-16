build:
	go build cmd


run: build
	./cmd/cmd


docker:
	docker build -t weather .