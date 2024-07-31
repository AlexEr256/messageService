package repositories

import (
	"fmt"
	"github.com/AlexEr256/messageService/dto"
	"github.com/jmoiron/sqlx"
	"time"
)

type IConsumerRepository interface {
	Add(c *dto.MessageRequest) (*dto.MessageRequest, error)
	Get() (*dto.AggregationResp, error)
}

type ConsumerRepository struct {
	Db *sqlx.DB
}

func NewConsumerRepository(db *sqlx.DB) IConsumerRepository {
	return &ConsumerRepository{Db: db}
}

func (r ConsumerRepository) Add(messageRequest *dto.MessageRequest) (*dto.MessageRequest, error) {
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

	result, err := r.Db.NamedExec(query, message)

	fmt.Println("Insert", result, err)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r ConsumerRepository) Get() (*dto.AggregationResp, error) {
	rows, err := r.Db.Queryx("SELECT COUNT(*) FROM messages")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var count int
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	resp := &dto.AggregationResp{
		Total: count,
	}

	return resp, nil
}
