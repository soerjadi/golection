test:
	@@go test -v -cover -race -covermode=atomic ./test/...

migration-up:
	@@cd migrations && goose postgres "user=surya dbname=golection sslmode=disable" up

migration-up-test:
	@@cd migrations && goose postgres "user=surya dbname=golection_test sslmode=disable" up

.PHONY: test migration-up migration-up-test
