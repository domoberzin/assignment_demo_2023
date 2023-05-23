package main

import (
	"context"
	"math/rand"
	"strings"
	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/service"
	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/models"
	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
)

// IMServiceImpl implements the last service interface defined in the IDL.
type IMServiceImpl struct{
	MessageService service.MessageService
}

// NewIMServiceImpl returns a new IMServiceImpl.
func NewIMServiceImpl(messageService service.MessageService) *IMServiceImpl {
	return &IMServiceImpl{
		MessageService: messageService,
	}
}

func (s *IMServiceImpl) Send(ctx context.Context, req *rpc.SendRequest) (*rpc.SendResponse, error) {
	resp := rpc.NewSendResponse()

	msg := req.GetMessage()
	newMsg := models.Message{}

	newMsg.Sender = req.GetMessage().GetSender()
	newMsg.Recipient = req.GetMessage().GetChat()
	newMsg.Text = req.GetMessage().GetText()

	err := s.MessageService.Save(newMsg)

	if err != nil {
		resp.Code, resp.Msg = 500, err.Error()
		return resp, nil
	}

	resp.Code, resp.Msg = 0, "success"
	return resp, nil
}

func (s *IMServiceImpl) Pull(ctx context.Context, req *rpc.PullRequest) (*rpc.PullResponse, error) {
	resp := rpc.NewPullResponse()

	query := req.GetChat()

	sender, recipient := strings.Split(query, ":")[0], strings.Split(query, ":")[1]

	messages, err := s.MessageService.FindAll(sender, recipient)


	var rpcMessages []*rpc.Message
	for _, msg := range *messages {
		rpcMessages = append(rpcMessages, &rpc.Message{
			Sender: msg.Sender,
			Chat: msg.Sender + ":" + msg.Recipient,
			Text: msg.Text,
		})
	}

	if err != nil {
		resp.Code, resp.Msg = 500, err.Error()
		return resp, nil
	}

	resp.Code, resp.Msg, resp.Messages = 0, "success", rpcMessages
	return resp, nil
}

func areYouLucky() (int32, string) {
	if rand.Int31n(2) == 1 {
		return 0, "success"
	} else {
		return 500, "oops"
	}
}