.PHONY: test cover cover-html

test:
	go test ./...

cover:
	go test -cover ./...

cover-html:
	go test -coverprofile=coverage.out ./...
	findstr /V -v "test/" coverage.out | findstr /V -v "internal/domain/" | findstr /V -v "cmd/" | findstr /V -v "internal/handler/routes/" > coverage.filtered
	go tool cover -html=coverage.filtered