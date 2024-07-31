package repositories

import (
	"github.com/AlexEr256/messageService/dto"
	"github.com/jmoiron/sqlx"
	"time"
)

type Message struct {
	Creator   string    `db:"creator"`
	Recipient string    `db:"recipient"`
	Mail      string    `db:"mail"`
	Created   time.Time `db:"created"`
}

type IProducerRepository interface {
	Add(c *dto.MessageRequest) (*dto.MessageResponse, error)
}

type ProducerRepository struct {
	Db *sqlx.DB
}

func NewProducerRepository(db *sqlx.DB) IProducerRepository {
	return &ProducerRepository{Db: db}
}

func (r ProducerRepository) Add(messageRequest *dto.MessageRequest) (*dto.MessageResponse, error) {
	message := &Message{
		Creator:   messageRequest.Creator,
		Recipient: messageRequest.Recipient,
		Mail:      messageRequest.Mail,
		Created:   time.Now(),
	}

	query := `INSERT INTO
				messages(creator, recipient, mail, created)
			VALUES
				(:creator, :recipient, :mail, :created);`

	_, err := r.Db.NamedExec(query, message)

	if err != nil {
		return nil, err
	}

	response := &dto.MessageResponse{Status: true}

	return response, nil
}
