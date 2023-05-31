package main

import (
	"context"
	"math/rand"
	"fmt"
	"strings"
	"time"
	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
)

// IMServiceImpl implements the last service interface defined in the IDL.
type IMServiceImpl struct{}

func (s *IMServiceImpl) Send(ctx context.Context, req *rpc.SendRequest) (*rpc.SendResponse, error) {
	
	timestamp := time.Now().UnixNano()

	roomID, err := getRoomID(req.Message.GetChat())

	print(roomID)

	if err != nil {
		return nil, err
	}

	message := &Message{
		Sender: req.Message.GetSender(),
		Message: req.Message.GetText(),
		Timestamp: timestamp,
	}

	err = db.SaveMessage(ctx, roomID, message)

	if err != nil {
		return nil, err
	}

	resp := rpc.NewSendResponse()
	resp.Code, resp.Msg = 0, "success"
	return resp, nil

}

func (s *IMServiceImpl) Pull(ctx context.Context, req *rpc.PullRequest) (*rpc.PullResponse, error) {

	roomID, err := getRoomID(req.GetChat())

	if err != nil {
		return nil, err
	}


	messages, err := db.GetMessages(ctx, roomID)

	finalMessages := make([]*rpc.Message, 0)

	if err != nil {
		return nil, err
	}

	for _, message := range messages {
		finalMessages = append(finalMessages, &rpc.Message{
			Chat:     req.GetChat(),
			Sender: message.Sender,
			Text: message.Message,
			SendTime: message.Timestamp,
		})

	}


	resp := rpc.NewPullResponse()
	resp.Messages = finalMessages
	resp.Code, resp.Msg = 0, "success"
	return resp, nil
}

func areYouLucky() (int32, string) {
	if rand.Int31n(2) == 1 {
		return 0, "success"
	} else {
		return 500, "oops"
	}
}

func getRoomID(chat string) (string, error) {

	participants := strings.Split(chat, ":")
	print(participants)

	if len(participants) != 2 {
		return "", fmt.Errorf("invalid chat: %s %s", participants, chat)
	}

	part1 := participants[0]
	part2 := participants[1]

	if part1 > part2 {
		return part1 + ":" + part2, nil
	} else {
		return part2 + ":" + part1, nil
	}


}
