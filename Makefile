MODULE=github.com/mkafonso/go-cloud-challenge
PKG=./__tests__/entity ./__tests__/usecase

.PHONY: test

test:
	go test $(PKG) -v
