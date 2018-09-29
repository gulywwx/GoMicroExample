package main

import (
	"GoMicroExample/service"
	"GoMicroExample/service/user/proto"
	userApi "GoMicroExample/service/user/proto"
	"context"
	"encoding/json"
	"github.com/micro/go-api/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"log"
)

type UserService struct {
}

func (us *UserService) Login(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	if req.Method != "POST" {
		return errors.BadRequest("go.micro.api.user", "require post")
	}

	ct, ok := req.Header["Content-Type"]
	if !ok || len(ct.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "need content-type")
	}

	if ct.Values[0] != "application/json" {
		return errors.BadRequest("go.micro.api.user", "expect application/json")
	}

	var userInfo user.UserInfo
	json.Unmarshal([]byte(req.Body), &userInfo)

	token, e := service.Encode(&userInfo)
	if e != nil {
		return e
	}
	b, _ := json.Marshal(map[string]string{
		"token": token,
	})
	rsp.Body = string(b)
	return nil
}

func main() {
	userService := micro.NewService(
		micro.Name("go.micro.api.user"),
		micro.WrapHandler(service.AuthWrapper),
	)

	userService.Init()

	userApi.RegisterUserHandler(userService.Server(), &UserService{})

	if err := userService.Run(); err != nil {
		log.Fatal(err)
	}
}