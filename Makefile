
GO111MODULE=on

build:
	cd reposam && go build -tags netgo \
		-ldflags '-w -extldflags "-static"'

install:
	install -m755 reposam/reposam /usr/local/bin

fixup:
	sed -i 's|CowYoSam|RepoSam|g' *.go
	sed -i 's|cowyosam|reposam|g' *.go
	sed -i 's|CowYoSam|RepoSam|g' reposam/*.go
	sed -i 's|cowyosam|reposam|g' reposam/*.go