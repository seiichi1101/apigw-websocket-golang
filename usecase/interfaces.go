package usecase

import (
	"lambda/domain/connection"
)

type IMessageSender interface {
	SendMessage(id connection.Id, message string) error
}

type IConnectionRepository interface {
	Find(id connection.Id) error
	Save(id connection.Id) error
	Delete(id connection.Id) error
}

type IUsecase interface {
	OnConnect(connectorId string) error
	OnDisconnect(connectorId string) error
	SendMessage(targetId string, message string) error
}
