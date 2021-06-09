package usecase

import "lambda/domain/connection"

func (u Usecase) SendMessage(targetId string, message string) error {
	id, err := connection.NewId(targetId)
	if err != nil {
		return err
	}
	if err := u.connectionRepository.Find(id); err != nil {
		return err
	}
	if err := u.messageSender.SendMessage(id, message); err != nil {
		return err
	}

	return nil
}
