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

func (r *RedisClient) GetMessages(ctx context.Context, roomID string, reverse bool, cursor int64, end int64, limit int64) ([]* Message, *bool, *int64, error) {

	var messages []*Message
	var members []string
	var err error

	key := roomID

	hasMore := false
	nextCursor := int64(0)
	
	var pHasMore *bool = &hasMore
	var pNextCursor *int64 = &nextCursor

	if reverse {
		members, err = r.client.ZRevRange(context.Background(), key, cursor, end).Result()
		if err != nil {
			return nil, pHasMore, pNextCursor, err
		}
	} else {
		members, err = r.client.ZRange(context.Background(), key, cursor, end).Result()
		if err != nil {
			return nil, pHasMore, pNextCursor, err
		}
	}



	// check number of messages, and return boolean for hasMore, as well as the next cursor

	counter := int64(0)

	for _, member := range members {
		if (counter >= limit) {
			hasMore = true
			nextCursor = end
			break
		}
		var message Message
		err := json.Unmarshal([]byte(member), &message)
		if err != nil {
			return nil, pHasMore, pNextCursor, err
		}
		messages = append(messages, &message)
		counter += 1
	}
	
	return messages, pHasMore, pNextCursor, nil
}

