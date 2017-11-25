.DEFAULT_GOAL := start

start:
	docker-compose build designer
	docker-compose up -d

stop:
	docker-compose down

restart: stop start

test:
	docker build -t planetary-gear-designer .
	docker run --rm planetary-gear-designer bash -c "echo 'test'"

.PHONY: test
