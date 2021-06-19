compile:
	echo "Compiling for every OS and Platform"
	GOOS=darwin GOARCH=amd64 go build -o ../bin/env-mac main.go
	GOOS=linux GOARCH=386 go build -o ../bin/env-linux main.go
	GOOS=linux GOARCH=arm go build -o ../bin/env-linux-arm main.go
	GOOS=windows GOARCH=386 go build -o ../bin/env-windows main.go