package commands

var Commands = map[string]func() string{
	"ECHO": handleEcho,
	"PING": handlePing,
}

func handleEcho() string {
	return "ECHO command has been called"
}

func handlePing() string {
	return "+PONG\r\n"
}
