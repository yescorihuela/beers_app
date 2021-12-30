package app

func Run() error {
	server := NewServer("localhost", 8081)
	return server.Run()
}