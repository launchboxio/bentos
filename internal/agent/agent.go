package agent

import (
	"fmt"
	"time"
)

func Run() {
	for {
		fmt.Println("Agent ping")
		time.Sleep(5 * time.Second)
	}

	// Load our adapter
	// Initiate the connection
	// Register our event handlers with the appropriate functions
}
