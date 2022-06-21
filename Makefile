.PHONY: build clean
output_folder = "app/"

build: main.go
	@go build -o ${output_folder} -trimpath -ldflags='-s -w'

clean:
	@rm -r ${output_folder}
