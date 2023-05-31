package main

import (
    "context"
    "encoding/json"
    "github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func (c *RedisClient) NewRedisClient() error {

	r := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	if err := r.Ping(context.Background()).Err(); err != nil {
		return err
	}

	c.client = r

	return nil
}

func (r *RedisClient) SaveMessage(ctx context.Context, roomID string, message *Message) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	key := roomID

	member := redis.Z{
		Score:  float64(message.Timestamp),
		Member: data,
	}

	return r.client.ZAdd(context.Background(), key, member).Err()
}

func (r *RedisClient) GetMessages(ctx context.Context, roomID string) ([]* Message, error) {

	var messages []*Message


	key := roomID

	members, err := r.client.ZRevRange(context.Background(), key, 0, -1).Result()

	if err != nil {
		return nil, err
	}

	for _, member := range members {
		var message Message
		err := json.Unmarshal([]byte(member), &message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}

	return messages, nil
}

