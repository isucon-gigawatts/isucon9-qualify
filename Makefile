export GO111MODULE=on

all: bin/benchmarker bin/benchmark-worker bin/payment bin/shipment

bin/benchmarker: cmd/bench/main.go bench/**/*.go
	go build -o bin/benchmarker cmd/bench/main.go

bin/benchmark-worker: cmd/bench-worker/main.go
	go build -o bin/benchmark-worker cmd/bench-worker/main.go

bin/payment: cmd/payment/main.go bench/server/*.go
	go build -o bin/payment cmd/payment/main.go

bin/shipment: cmd/shipment/main.go bench/server/*.go
	go build -o bin/shipment cmd/shipment/main.go

vet:
	go vet ./...

errcheck:
	errcheck ./...

staticcheck:
	staticcheck -checks="all,-ST1000" ./...

clean:
	rm -rf bin/*

## benchmerker
.PHONY: bench
bench:
	./bin/benchmarker --target-url=http://127.0.0.1:80

## Nginx
load-nginx-conf:
	sudo cp config/etc/nginx/nginx.conf /etc/nginx/nginx.conf
	sudo systemctl restart nginx.service

format-access-log:
	@cat /var/log/nginx/access.log | alp ltsv --sort=max -r

## MySQL
load-mysqld-conf:
	sudo cp config/etc/mysql/my.cnf /etc/mysql/my.cnf
	sudo systemctl restart mysql.service

enable-slow-query:

dump-slow-query:
	sudo mysqldumpslow -s t /var/log/mysql/mysql-slow.log
