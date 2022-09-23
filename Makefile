build-docker:
	@docker build -t data-wilayah .
run-docker:
	@docker container run -ti --init --rm --name data-wilayah --memory="1g" --cpus="1" --net host data-wilayah:latest
run:
	@go run .