package service

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/backend/model"
)

type LoginService struct{}

func (login *LoginService) LoginRequest(request model.Login) {
	fmt.Println(request.Command)
}
