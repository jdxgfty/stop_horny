.PHONY: build clean docker build-docker
output_folder = "app/"

build: main.go
	@go build -o ${output_folder} -trimpath -ldflags='-s -w'

build-docker: build
	docker build -t "stop_horny" .

clean:
	@rm -r ${output_folder}
