default: all
all: 
	npm run docs:dev

.phony: build
build:
	npm run docs:build

.PHONY: publish
publish:
	git add .
	git commit -m "update"
	git push
	git checkout websites
	git pull
	git merge master
	rm -rf ./websites/*
	npm run docs:build
	git add .
	git commit -m "update"
	git push

