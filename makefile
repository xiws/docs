default: all
all: 
	npm run docs:dev

.phony: build
build:
	npm run docs:build

.PHONY: publish
publish:
	npm run docs:build
	git add .
	git commit -m "update"
	git push
	
