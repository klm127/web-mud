package workers

func StartAllWorkers() {
	go StartIdleConnectionCleaner()
	go StartDirtyRoomClean(1)
	go StartMoveProcesser(0.25)
}
