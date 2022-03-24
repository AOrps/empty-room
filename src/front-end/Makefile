SHELL=/bin/bash
OUTPUT= main.o

all: compile

compile:
	@go build -o $(OUTPUT) main.go

run: compile
	@./$(OUTPUT)

clean:
ifneq (,$(wildcard ./$(OUTPUT)))
	@rm ./$(OUTPUT)
endif