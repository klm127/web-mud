package workers

func StartAllWorkers() {
	go StartIdleConnectionCleaner()
	go StartDirtyRoomClean(1)
}
