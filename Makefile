run:
	@go run app/main.go

watch:
	@air

docker-push:
	@docker build -t subrotokumar/rover .
	@docker push subrotokumar/rover