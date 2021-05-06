package v1

import (
	application "ddd_demo/internal/usecase/service"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/codes"
	"strconv"
)

type UserServiceHandler struct {
	uc *application.UserUseCase
}

func NewUserServiceHandler(uc *application.UserUseCase, router *gin.Engine) *gin.Engine {
	uh := new(UserServiceHandler)
	uh.uc = uc
	router.GET("/v1/users/:user_id", uh.GetUser)
	return router
}

func (uh *UserServiceHandler) GetUser(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.Param("user_id"), 10, 0)
	if userID <= 0 {
		c.JSON(runtime.HTTPStatusFromCode(codes.NotFound),
			map[string]interface{}{
				"code":    codes.NotFound,
				"message": codes.NotFound.String(),
			})
		return
	}
	user, err := uh.uc.GetUser(c, userID)
	if err != nil {
		c.JSON(runtime.HTTPStatusFromCode(codes.NotFound),
			map[string]interface{}{
				"code":    codes.NotFound,
				"message": codes.NotFound.String(),
				"reason":  err.Error(),
			})
		return
	}
	c.JSON(200, &GetUserResponse{
		UserID:    user.UserID,
		Nickname:  user.Nickname,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})

}
