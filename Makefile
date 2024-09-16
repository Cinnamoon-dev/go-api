build:
	rm *.db &> /dev/null
	
run: build
	go run main.go