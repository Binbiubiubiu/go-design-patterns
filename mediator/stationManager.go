package main

type StationManager struct {
	isPlatFormFree bool
	trainQueue     []Train
}

func newStationManger() *StationManager {
	return &StationManager{
		isPlatFormFree: true,
	}
}

func (s *StationManager) canArrive(t Train) bool {
	if s.isPlatFormFree {
		s.isPlatFormFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) notifyAboutDeparture() {
	if !s.isPlatFormFree {
		s.isPlatFormFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}
