export GO111MODULE=on
export GOOGLE_CLOUD_PROJECT=isucon-gigawatts
export ENABLE_TRACE=true
export ENABLE_PROFILE=true
export GOOGLE_APPLICATION_CREDENTIALS=/home/ikedamorito/isucon-gigawatts-06b12dcfbd07.json

CAMPAIGN=0

.PHONY: build
build: *.go
	@sed -i -e 's/Campaign: [01234]/Campaign: ${CAMPAIGN}/g' main.go
	@echo "### ENVIOROMENT INFO ###"
	@echo CAMPAIGN = ${CAMPAIGN}
	@go build -o isucari
	@sed -i -e 's/Campaign: [01234]/Campaign: 0/g' main.go

.PHONY: run
run: build
	@./isucari

.PHONY: vendor
vendor:
	go mod vendor
