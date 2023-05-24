package HotRemote

import (
	"strconv"
)

//go:generate go run gen.go

type Session struct {
	client    *Client
	targetBox string
}

func (c *Client) CreateSession(targetBox string) *Session {
	return &Session{
		client:    c,
		targetBox: targetBox,
	}
}

func (s *Session) ButtonPress(buttonID int) error {
	return s.client.ButtonPress(s.targetBox, buttonID)
}

func (s *Session) SwitchToChannel(channelNumber int) {
	s.SendStr(strconv.Itoa(channelNumber))
}

func StrToAssci(value string) []int {
	runes := []rune(value)

	var result []int

	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}

	return result
}

func (s *Session) SendStr(value string) {
	for _, v := range StrToAssci(value) {
		if err := s.ButtonPress(v); err != nil {
			return
		}
	}
}
