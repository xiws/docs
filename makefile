
SyncResult = $(shell  DirStructEcho -path="./Skill");


default: all
all: 
	npm run docs:dev

.phony: build
build:
	npm run docs:build

.phony: push
push:
	git add .
	git commit -m "update"
	git push

.PHONY: Sync
Sync:
	echo 'var list=$(SyncResult);' > .vitepress/sidebarConfig.js
	echo "export const sidebarConfig = list;" >> .vitepress/sidebarConfig.js

