run:
	export ARDUINO_DEVICE=/dev/tty.usbmodem14101 &&\
	export COLLECTOR_API=http://localhost:8080/value &&\
	export TERMINAL=T01 &&\
	export SENSOR=temperature &&\
	go run .

rasperrypi:
	GOOS=linux GOARCH=arm GOARM=5 go build -o sensor_collector