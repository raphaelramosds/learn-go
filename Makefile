.PHONY: heranca excecoes

run-%:
	@go run $*/main.go

test-%:
	@cd $* && go test