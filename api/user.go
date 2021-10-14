package api

import (
	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/core"

	"github.com/gofiber/fiber/v2"
)

func Hello(ctx *fiber.Ctx) error {
	return ctx.JSON("Hello")
}
func Login(c *fiber.Ctx) error {
	var req request.Login
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}

	data, err := biz.Login(c.Context(), req.Username, req.Password)
	if err != nil {

		return err
	}

	return core.SetData(c, data)

}

func Register(c *fiber.Ctx) error {
	var req request.Register
	if err := core.ParseRequest(c, &req); err != nil {

		return err
	}

	err := biz.Register(c.Context(), req)
	if err != nil {
		return err
	}
	return core.SuccessResp(c)
}

func UpdateMe(c *fiber.Ctx) error {
	var req request.UpdateMe
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}
	user := biz.LocalUser(c)
	err := biz.UpdateMe(c.Context(), req, user.ID)
	if err != nil {
		return err
	}
	return core.SuccessResp(c)
}

func ChangeMyPassword(c *fiber.Ctx) error {

	var req request.ChangeMyPassword
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}

	var uid int

	err := biz.ChangeMyPassword(c.Context(), req, uid)
	if err != nil {

		return err
	}
	return core.SuccessResp(c)

}

func GetMyPermissions(c *fiber.Ctx) error {
	user := biz.LocalUser(c)

	data, err := biz.GetMyPermissions(c.Context(), user.ID)
	if err != nil {

		return err
	}
	core.SetData(c, data)
	return nil
}

func GetMyInfomation(c *fiber.Ctx) error {
	//uid, err := uid(c)
	//if err != nil {
	//
	//	return err
	//}
	//data, err := biz.GetMyInfomation(uid)
	//if err != nil {
	//
	//	return err
	//}
	//return core.SetData(c, data)
	return nil
}

//TODO 对cms意义并不大，先不实现
func RefreshToken(c *fiber.Ctx) error {
	return nil
}
