
GO111MODULE=on

build: fmt readme
	go build -tags netgo \
		-ldflags '-w -extldflags "-static"' \
		-o reposam/reposam ./reposam

try: build
	cd tmp && ../reposam/reposam

readme:
	cat .readme.md > README.md
	@echo '' | tee -a README.md
	@echo '' | tee -a README.md
	@echo '' | tee -a README.md
	./reposam/reposam -h 2>&1 | sed 's|  |         |g' | sed 's|                	|            |g' | sed 's|Usage|        Usage|g' |tee -a README.md

install:
	install -m755 reposam/reposam /usr/local/bin

fmt:
	gofmt -w *.go */*.go

fixup:
	sed -i 's|CowYoSam|RepoSam|g' *.go
	sed -i 's|cowyosam|reposam|g' *.go
	sed -i 's|CowYoSam|RepoSam|g' reposam/*.go
	sed -i 's|cowyosam|reposam|g' reposam/*.go