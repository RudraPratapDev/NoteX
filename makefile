build:
	@ go build -o notex .

run:
	@ ./notex

start: build run