

all: build

build:
	cd dockbroker; go install
	cd docksubmit; go install
