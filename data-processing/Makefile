SHELL=/bin/bash

all = json

json: 
	@python analyze.py | python3 -m json.tool > out.json

clean:
ifneq (,$(wildcard ./out.json))
	@rm ./out.json
endif