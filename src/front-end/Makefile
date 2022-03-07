SHELL=/bin/bash
OUTPUT= main.o

all: run

compile:
	@go build -o $(OUTPUT) main.go

run: compile
	@./$(OUTPUT)

clean:
ifneq (,$(wildcard ./$(OUTPUT)))
	@rm ./$(OUTPUT)
endif