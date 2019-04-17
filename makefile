GOPATH = $(PWD)/build
export GOPATH
URL = github.com/chenji-kael
REPO = go-bootstrap
URLPATH = $(PWD)/build/src/$(URL)

build:
	[ -d $(URLPATH) ] || mkdir -p $(URLPATH)
	ln -nsf $(PWD) $(URLPATH)/$(REPO)
	go build -o go-bootstrap $(URLPATH)/$(REPO)/src/main.go
