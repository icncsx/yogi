gen:
	kitex -module github.com/cloudwego/kitex-examples ./echo.thrift

run:
	go run .

stop:
	kill $(lsof -t -i:8888)