export GO111MODULE=on
export GOOGLE_CLOUD_PROJECT=isucon-gigawatts
export ENABLE_TRACE=false
export ENABLE_PROFILE=false
export GOOGLE_APPLICATION_CREDENTIALS=/home/ikedamorito/isucon-gigawatts-06b12dcfbd07.json

CAMPAIGN=0

all: isucari

rebuild: *.go
	@sed -i -e 's/Campaign: [01234]/Campaign: ${CAMPAIGN}/g' main.go
	@echo CAMPAIGN = ${CAMPAIGN}
	@go build -o isucari
	@sed -i -e 's/Campaign: [01234]/Campaign: 0/g' main.go

run: rebuild
	@./isucari

vendor:
	go mod vendor
