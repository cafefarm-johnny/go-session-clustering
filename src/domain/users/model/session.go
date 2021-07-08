package model

type Session struct {
	UUID     string
	Username string
}

func NewSession(uuid, username string) *Session {
	return &Session{
		UUID:     uuid,
		Username: username,
	}
}
