# 学习go语言日记

## 1. hello world
cd ./hello
go mod init example.com/hello
go run .

## 2. call you code from another module

- In Go, code executed as an application must be in a main package.
- go mod edit -replace example.com/greetings=../greetings
- go mod tidy