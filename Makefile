.PHONY:	test

test:
	@cd app/http/controllers/api && go test -v
