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
}
