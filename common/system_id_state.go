package common

type SystemIDState struct {
	systemID string
}


func (s *SystemIDState) GetSystemID() string {
	return s.systemID
}
