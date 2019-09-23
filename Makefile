build:
	go build -o scout_regiment main.go

dep:
	dep ensure

launch:
	make build && make run

run:
	./scout_regiment
