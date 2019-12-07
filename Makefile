

GO111MODULE=on
VERSION=0.0.1
packagename=reposam
USER_GH=eyedeekay

build: fmt readme
	go build -tags netgo \
		-ldflags '-w -extldflags "-static"' \
		-o reposam/reposam ./reposam

try: build
	cd tmp && ../reposam/reposam

readme:
	sed 's|See README.md||g' .readme.md > README.md
	@echo '' | tee -a README.md
	./reposam/reposam -h 2>&1 | \
		sed 's|  |         |g' | \
		sed 's|                	|            |g' | \
		sed 's|Usage|        Usage|g' | \
		tee -a README.md

install:
	install -m755 reposam/reposam /usr/local/bin

fmt:
	gofmt -w *.go */*.go

fixup:
	sed -i 's|CowYoSam|RepoSam|g' *.go
	sed -i 's|cowyosam|reposam|g' *.go
	sed -i 's|CowYoSam|RepoSam|g' reposam/*.go
	sed -i 's|cowyosam|reposam|g' reposam/*.go

tag: build
	cat .readme.md | gothub release -s $(GITHUB_TOKEN) -u $(USER_GH) -r $(packagename) -t v$(VERSION) -d -

#del:
#	gothub delete -s $(GITHUB_TOKEN) -u $(USER_GH) -r $(packagename) -t v$(VERSION)