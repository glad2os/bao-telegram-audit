.PHONY: build

build:
	docker buildx build --platform linux/amd64 -t glad2os/bao-telegram-audit:latest .

push:
	docker push glad2os/bao-telegram-audit:latest