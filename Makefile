all:
	go build -o bookings-server.exe ./cmd/web
	./bookings-server
