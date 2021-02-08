
ROOT := cmd/certmgr

debug:
	cd $(ROOT) && go run .

build:
	# cd $(ROOT) && go build -o ../../bin/certmgr .
	cd $(ROOT) && GOOS=linux go build -o ../../bin/certmgr-linux-amd64 .
