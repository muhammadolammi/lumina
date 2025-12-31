package main

func dispatcher(input <-chan string, fans []chan<- string) {
	for msg := range input {
		for _, fan := range fans {
			// We send the same log to every 'fan' channel
			fan <- msg
		}
	}
	// When input is done, close all fans
	for _, fan := range fans {
		close(fan)
	}
}
