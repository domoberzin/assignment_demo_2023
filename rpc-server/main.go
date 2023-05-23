package main

import (
	"log"
	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/infrastructure"
	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/service"
	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/repository"
	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/models"


	rpc "github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc/imservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"etcd:2379"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	infrastructure.LoadEnv()  //loading env
	db := infrastructure.NewDatabase()
	messageRepository := repository.NewMessageRepository(db)
	messageService := service.NewMessageService(messageRepository)
	db.DB.AutoMigrate(&models.Message{})

	svr := rpc.NewServer(NewIMServiceImpl(messageService), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "demo.rpc.server",
	}))

	println("rpc server start")



	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
