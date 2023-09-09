package workers

func StartAllWorkers() {
	go StartIdleConnectionCleaner()
}
