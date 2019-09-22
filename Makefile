build:
	go build -o scout_regiment main.go

dep:
	dep ensure

run:
	./scout_regiment
