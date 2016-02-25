

BINARY=gobeepme



all: clean build ui install

build:
	@echo "\n*** Building gobeepme ***\n"
	go build -o ${BINARY} gobeepme.go

.PHONY: install
install:
	@echo "*** Installing gobeepme ***\n"
	go install ./...

.PHONY: clean
clean:
	@echo "\n*** Cleaning gobeepme ***\n"
	cd static/app && find . -name "*.js.map" -type f -delete
	cd static/app && find . -name "*.js" -type f -delete
	#cd static && rm -rf node_modules
	find . -name "*.log" -type f -delete
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: ui
ui:
	@echo "\n*** Intalling npm dependencies for gobeepme ***\n"
	cd static && npm install
	@echo "\n*** Compiling TypeScript **\n"
	cd static && npm run tsc
