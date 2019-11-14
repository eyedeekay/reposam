
GO111MODULE=on

build:
	go build -tags netgo \
		-ldflags '-w -extldflags "-static"' \
		-o reposam/reposam

install:
	install -m755 reposam/reposam /usr/local/bin

fmt:
	gofmt -w *.go */*.go

fixup:
	sed -i 's|CowYoSam|RepoSam|g' *.go
	sed -i 's|cowyosam|reposam|g' *.go
	sed -i 's|CowYoSam|RepoSam|g' reposam/*.go
	sed -i 's|cowyosam|reposam|g' reposam/*.go