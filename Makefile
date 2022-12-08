# start user service
start-us:
	go run user/cmd/server/main.go

# start product service
start-ps:
	go run product/cmd/server/main.go 

# start gateway service
start-g:
	go run gateway/cmd/web/main.go 