package service

import (
    "github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/repository"
    "github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/models"
)

//MessageService MessageService struct
type MessageService struct {
    repository repository.MessageRepository
}

//NewMessageService : returns the MessageService struct instance
func NewMessageService(r repository.MessageRepository) MessageService {
    return MessageService{
        repository: r,
    }
}

//Save -> calls message repository save method
func (p MessageService) Save(message models.Message) error {
    return p.repository.Save(message)
}

//FindAll -> calls message repo find all method
func (p MessageService) FindAll(sender string, recipient string) (*[]models.Message, error) {
    return p.repository.FindAll(sender, recipient)
}