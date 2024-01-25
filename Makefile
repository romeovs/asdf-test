
asdf-plugins:
	@echo "Installing asdf plugins"
	asdf plugin add golang https://github.com/asdf-community/asdf-golang.git
	asdf plugin add nodejs https://github.com/asdf-vm/asdf-nodejs.git

install:
	asdf install

run-go:
	go run ./test.go

run-node:
	node ./test.js
