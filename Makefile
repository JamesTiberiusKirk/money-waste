tsgen: 
	npm run build

tstypegen:
	tygo generate

buildts: 
	npm run build

dev_setup:
	go install github.com/JamesTiberiusKirk/tygo@v0.2.5
	go install github.com/cespare/reflex@latest

install:
	go get ./...
	go mod vendor

lint:
	golangci-lint run

lint-less:
	golangci-lint run --color always | less -R

