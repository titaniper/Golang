https://selenehyun.notion.site/Golang-Study-e856d5b0c8434119b4a574c812459b6f

# 빌드
go build

# 실행
go run main.go

# 설정
```azure
export GOPATH= ~ >> ~/.zshrc
export PATH=$PATH:$GOPATH/bin >> ~/.zshrc
source ~/.zshrc
```

# 새로운 패키지 생성
```
echo $GOPATH
mkdir -p $GOPATH/src/github.com/titaniper/webserver
cd $GOPATH/src/github.com/titaniper/webserver
go mod init github.com/titaniper/webserver
code $GOPATH/src/github.com/titaniper/webserver
```

# 패키지
```
go build .
go run main.go
go get -u github.com/gorilla/mux
```