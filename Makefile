.PHONY: tools
tools:
	$(MAKE) -C cmd/troptgen install

.PHONY: gen
gen: tools
	rm options_gen.go
	go generate ./...
	find ./cache -maxdepth 1 ! -path ./cache -type d | xargs -I % sh -c 'cd %; go generate ./...'

.PHONY: test
test: tools
	go test ./...
	find ./cache -maxdepth 1 ! -path ./cache -type d | xargs -I % sh -c 'cd %; go test ./...'

.PHONY: update-dep-version
update-dep-version:
	test -n "$(TAG)"  # $$TAG
	sh -c 'cd mocks; go get github.com/RangelReale/trcache@$(TAG); go mod tidy'
	find ./cache -maxdepth 1 ! -path ./cache -type d | xargs -I % sh -c 'cd %; go get github.com/RangelReale/trcache@$(TAG); go get github.com/RangelReale/trcache/mocks@$(TAG); go mod tidy'

git-status:
	@status=$$(git status --porcelain); \
	if [ ! -z "$${status}" ]; \
	then \
		echo "Error - working directory is dirty. Commit those changes!"; \
		exit 1; \
	fi

.PHONY: gittag
gittag: git-status update-dep-version
	test -n "$(TAG)"  # $$TAG
	git commit -a -m "Release $(TAG)"
	sh -c 'git tag mocks/$(TAG)'
	find cache -maxdepth 1 ! -path cache -type d | xargs -I % sh -c 'git tag %/$(TAG)'
	git tag $(TAG)
