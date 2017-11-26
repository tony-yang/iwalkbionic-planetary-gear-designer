.DEFAULT_GOAL := start

start:
	docker build -t planetary-gear-designer .
	docker run -itd --rm planetary-gear-designer bash

stop:
	docker container ls | grep planetary-gear-designer | awk '{print $$1}' | xargs docker container stop

restart: stop start

test:
	docker build -t planetary-gear-designer .
	docker run --rm planetary-gear-designer bash -c "echo 'test'"

.PHONY: test
