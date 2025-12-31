package main

func produceLogs(source LogSource, logsChan chan<- string) {
	for {
		msg, ok := source.GetNextLog()
		if !ok {
			break
		}
		logsChan <- msg
	}
}
