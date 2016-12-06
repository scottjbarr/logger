package main

import "github.com/scottjbarr/logger"

func main() {
	log := logger.New("SomeProcess")

	log.Debug("Hello %v, nice to see you %s", 42, "today")
	log.Info("Yar %s", "hello")
	log.Warn("Yar %s", "hello")
	log.Error("Yar %s", "hello")
	log.Info("Boring")
}
