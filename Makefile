export GOOS=linux

.PHONY: clean
clean:
	@rm -rf ./bin/*

.PHONY: build
build:
	@ls handler | xargs -i go build -o bin/{} handler/{}/main.go

.PHONY: deploy
deploy: clean build
	@rm -f terraform/function.zip
	@cd terraform && terraform apply

.PHONY: destroy
destroy:
	@cd terraform && terraform destroy