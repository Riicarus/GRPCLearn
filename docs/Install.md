# Protobuf 安装
1. 在 protobuf 官网下载最新版 适用于 win64 的 protobuf.
2. 解压文件夹, 为 protoc.exe 添加环境变量: `D:\software\protoc\bin\protoc.exe`.
3. 在命令行输入 `protoc --version` 来测试是否安装成功.
4. vscode 安装 vscode-proto3 和 clang-format 插件, 设置配置: 
 ```json
"protoc": {
    "path": "D:\\software\\protoc\\bin\\protoc.exe",
 }
```
5. 安装适配 go 的 protobuf 插件: protoc-gen-go, protoc-gen-go-grpc.  
```bash
go get github.com/golang/protobuf/protoc-gen-go
go get github.com/golang/protobuf/protoc-gen-go-grpc
```
6. 安装好后, 切换到 `{GOPATH}\pkg\mod\github\golang` 下, 找到 protobuf 对应的包, 进入其中并且运行 `go install .\proto-gen-go`. 再切换到 `{GOPATH}\pkg\mod\google.golang.org\grpc\cmd` 下, 找到 protoc-gen-go-grpc 对应的包, 运行 `go install .\protoc-gen-go-grpc`.
7. 切换到 `{GOPATH}\bin` 下, 将 protoc-gen-go.exe 和 protoc-gen-go-grpc.exe 复制到 `{GOROOT}\bin` 中.