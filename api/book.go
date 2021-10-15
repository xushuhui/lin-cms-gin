package api

import (
	"github.com/gofiber/fiber/v2"
	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/utils"
)

func GetBooks(c *fiber.Ctx) error {
	// TODO book 接口少了权限判断

	size := core.GetSize(c)
	page := core.GetPage(c)
	data, err := biz.GetBookAll(c.Context(), page, size)
	if err != nil {
		return err
	}

	total, err := biz.GetBookTotal(c.Context())
	if err != nil {
		return err
	}

	core.SetPage(c, data, total)

	return nil
}

func UpdateBook(c *fiber.Ctx) error {
	var req request.UpdateBook
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}
	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return err
	}

	err = biz.UpdateBook(c.Context(), id, req)
	if err != nil {
		return err
	}
	return core.SuccessResp(c)
}

func CreateBook(c *fiber.Ctx) error {
	var req request.CreateBook
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}
	err := biz.CreateBook(c.Context(), req)
	if err != nil {
		return err
	}
	return core.SuccessResp(c)
}

func DeleteBook(c *fiber.Ctx) error {
	// TODO 没有软删除，ent默认没有软删除待加强
	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return err
	}
	err = biz.DeleteBook(c.Context(), id)
	if err != nil {
		return err
	}
	return core.SuccessResp(c)
}

func GetBook(c *fiber.Ctx) error {
	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return err
	}
	data, err := biz.GetBook(c.Context(), id)
	if err != nil {
		return err
	}
	core.SetData(c, data)
	return nil
}