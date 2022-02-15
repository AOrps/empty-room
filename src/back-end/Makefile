SHELL=/bin/bash

all = compile

compile:
	@go build main.go

clean:
ifneq (,$(wildcard ./main))
	@rm ./main
endif