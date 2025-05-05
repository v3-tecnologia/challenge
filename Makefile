MODULE=github.com/mkafonso/go-cloud-challenge
PKG=./__tests__/entity

.PHONY: test

test:
	go test $(PKG) -v
