build:
	cd api &&SET GOOS=linux&&SET GOARCH=amd64&&go build -o micro ./main.go ./plugin.go
	cd service/greeter && protoc --proto_path=proto --proto_path=${GOPATH} --micro_out=proto --go_out=proto greeter.proto &&SET GOOS=linux&&SET GOARCH=amd64&&go build
	cd service/user && protoc --proto_path=proto --proto_path=${GOPATH} --micro_out=proto --go_out=proto user.proto &&SET GOOS=linux&&SET GOARCH=amd64&&go build
	cd config && gradlew build