package connection

import "fmt"

type Id string

func NewId(value string) (Id, error) {
	if len(value) < 1 {
		return "", fmt.Errorf("id must be at least one character")
	}
	id := Id(value)
	return id, nil
}

func (i Id) String() string {
	return string(i)
}
