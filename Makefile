.PHONY: build clean

build:
	go build -o img-processor

clean:
	rm -f img-processor out/*