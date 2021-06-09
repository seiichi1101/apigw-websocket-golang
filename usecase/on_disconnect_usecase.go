package usecase

import "lambda/domain/connection"

func (u Usecase) OnDisconnect(connectorId string) error {
	id, err := connection.NewId(connectorId)
	if err != nil {
		return err
	}
	if err := u.connectionRepository.Delete(id); err != nil {
		return err
	}
	return nil
}
