package responses

import (
	"fmt"
	"root/destinyhome/bungie/responses/destiny2"
)

type OuterRes struct {
	Response        innerRes
	ErrorCode       ErrorCode `json:"ErrorCode"`
	ThrottleSeconds int       `json:"ThrottleSeconds"`
	ErrorStatus     string    `json:"ErrorStatus"`
	Message         string    `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}

func (os *OuterRes) Error() string {
	return fmt.Sprintf("%#v", os)
}

type innerRes interface {
	Get() *destiny2.Character
}
