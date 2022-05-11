package user

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"

	"github.com/xiaomudk/kube-ybuild/internal/ecode"
	"github.com/xiaomudk/kube-ybuild/internal/repository"
	"github.com/xiaomudk/kube-ybuild/internal/service"
	"github.com/xiaomudk/kube-ybuild/pkg/errcode"
	"github.com/xiaomudk/kube-ybuild/pkg/logs"
)

// Get 获取用户信息
// @Summary 通过用户id获取用户信息
// @Description Get an user by user id
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param id path string true "用户id"
// @Success 200 {object} model.UserInfo "用户信息"
// @Router /users/:id [get]
func Get(c echo.Context) error {
	logs.Info("Get function called.")

	userID := cast.ToUint64(c.Param("id"))
	if userID == 0 {
		response.Error(c, errcode.ErrInvalidParam)
		return nil
	}

	// Get the user by the `user_id` from the database.
	u, err := service.Svc.Users().GetUserByID(c.Request().Context(), userID)
	if errors.Is(err, repository.ErrNotFound) {
		logs.Errorf("get user info err: %+v", err)
		response.Error(c, ecode.ErrUserNotFound)
		return nil
	}
	if err != nil {
		response.Error(c, errcode.ErrInternalServer.WithDetails(err.Error()))
		return nil
	}

	response.Success(c, u)
	return nil
}
