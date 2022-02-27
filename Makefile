build:
	go build -o bin/main.exe

test: 
	go test cli_converter/tests -v

run: 
	go run main.go

launch_cli_convert:
	./bin/main.exe convert --decimalToBinary 23


launch_cli_addition:
	./bin/main.exe addition --octalAdd 23,21

launch_cli_substraction:
	./bin/main.exe substraction --binarySubs 1111,1100
