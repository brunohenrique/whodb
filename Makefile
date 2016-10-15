PACKAGES=$(shell find . -name '*.go' -print0 | xargs -0 -n1 dirname | sort --unique)

metalint:
	gometalinter $(PACKAGES)

metalint-ci:
	gometalinter --disable-all \
		--enable=errcheck \
		--enable=gocyclo \
		--enable=deadcode \
		--enable=aligncheck \
		--enable=defercheck \
		--enable=structcheck \
		--enable=golint $(PACKAGES)

test:
	go test -v -race ./...

tools:
	go get -u github.com/alecthomas/gometalinter gopkg.in/redis.v5
	gometalinter --install
