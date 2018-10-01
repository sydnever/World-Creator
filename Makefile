LAUNCHER=./bin/launcher
SERVER=./bin/server
TARGETS=./targets

all: windows

linux:
	go build -o $(TARGETS)/launcher $(LAUNCHER)
	go build -o $(TARGETS)/server $(SERVER)

windows:
	xgo --targets=windows/amd64 -dest=$(TARGETS) $(SERVER) 
	xgo --targets=windows/amd64 -dest=$(TARGETS) $(LAUNCHER) 
