package repository
import (
    "github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/models"
    "github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/infrastructure"
)

//MessageRepository -> MessageRepository
type MessageRepository struct {
    db infrastructure.Database
}

func NewMessageRepository(db infrastructure.Database) MessageRepository {
    return MessageRepository{
        db: db,
    }
}

func (p MessageRepository) Save(message models.Message) error {
    return p.db.DB.Create(&message).Error
}

func (p MessageRepository) FindAll(sender string, recipient string) (*[]models.Message, error) {
    var messages []models.Message
    var totalRows int64 = 0

    // queryBuider := p.db.DB.Order("created_at desc").Model(&models.Message{})

    // Search for messages where message matches sender and recipient
	queryBuider := p.db.DB.Model(&models.Message{})

	queryBuider = queryBuider.Where("message.sender = ? AND message.recipient = ?", sender, recipient)

	err := queryBuider.
		Find(&messages).
		Count(&totalRows).Error
    return &messages, err
}