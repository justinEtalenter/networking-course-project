package v1

import (
	"strconv"
	"yak/backend/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func (apiVX *ApiV1) registerBoardsHandlers(router fiber.Router) {
	group := router.Group("/projects/:pid/boards")
	// group.Get("/", apiVX.getBoards)
	group.Post("/", apiVX.createBoard)
	group.Get("/:bid", apiVX.getBoard)
	group.Put("/:bid", apiVX.updateBoard)
	group.Delete("/:bid", apiVX.deleteBoard)
}

// func (apiVX *ApiV1) getBoards(ctx *fiber.Ctx) error {
// 	response := &models.ApiResponse{}
// 	userId, err := apiVX.getUserId(ctx)
// 	if err != nil {
// 		response.Error(fiber.StatusInternalServerError, err.Error())
// 		return Send(ctx, response)
// 	}

// 	projectId := ctx.Params("pid")
// 	if projectId == "" {
// 		response.Error(fiber.StatusBadRequest, "Empty projectId")
// 		return Send(ctx, response)
// 	}

// 	response := apiVX.services.Board.GetAll(userId, projectId)
// 	return Send(ctx, response)
// }

func (apiVX *ApiV1) getBoard(ctx *fiber.Ctx) error {
	response := &models.ApiResponse{}
	userId, err := getUserId(ctx)
	if err != nil {
		response.Error(fiber.StatusInternalServerError, err.Error())
		return Send(ctx, response)
	}

	projectId, err := strconv.Atoi(ctx.Params("pid"))
	if err != nil || projectId == 0 {
		response.Error(fiber.StatusBadRequest, "Empty projectId")
		return Send(ctx, response)
	}

	boardId, err := strconv.Atoi(ctx.Params("bid"))
	if err != nil || boardId == 0 {
		response.Error(fiber.StatusBadRequest, "Empty boardId")
		return Send(ctx, response)
	}

	response = apiVX.services.Board.GetById(userId, projectId, boardId)
	return Send(ctx, response)
}

func (apiVX *ApiV1) createBoard(ctx *fiber.Ctx) error {
	response := &models.ApiResponse{}
	userId, err := getUserId(ctx)
	if err != nil {
		response.Error(fiber.StatusInternalServerError, err.Error())
		return Send(ctx, response)
	}

	projectId, err := strconv.Atoi(ctx.Params("pid"))
	if err != nil || projectId == 0 {
		response.Error(fiber.StatusBadRequest, "Empty projectId")
		return Send(ctx, response)
	}

	input := &models.Board{}
	if err := ctx.BodyParser(input); err != nil {
		response.Error(fiber.StatusBadRequest, err.Error())
		return Send(ctx, response)
	}

	response = apiVX.services.Board.Create(userId, projectId, input)
	return Send(ctx, response)
}

func (apiVX *ApiV1) updateBoard(ctx *fiber.Ctx) error {
	implementMe()
	board := models.Board{}
	return ctx.JSON(board)
}

func (apiVX *ApiV1) deleteBoard(ctx *fiber.Ctx) error {
	implementMe()
	board := models.Board{}
	return ctx.JSON(board)
}
