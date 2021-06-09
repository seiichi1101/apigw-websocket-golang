package usecase

type Usecase struct {
	connectionRepository IConnectionRepository
	messageSender        IMessageSender
}

func NewUsecase(connectionRepository IConnectionRepository, messageSender IMessageSender) *Usecase {
	return &Usecase{connectionRepository, messageSender}
}
