

all: build

build:
	cd api; go install
	cd dockbroker; go install
	cd docksubmit; go install
