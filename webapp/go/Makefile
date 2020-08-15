export GO111MODULE=on
CAMPAIGN=0

all: isucari

rebuild: *.go
	@sed -i -e 's/Campaign: [01234]/Campaign: ${CAMPAIGN}/g' main.go
	@echo CAMPAIGN=${CAMPAIGN}
	go build -o isucari
	@sed -i -e 's/Campaign: [01234]/Campaign: 0/g' main.go

run: rebuild
	./isucari
