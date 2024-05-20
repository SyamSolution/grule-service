package handler

import (
	"github.com/SyamSolution/grule-service/config"
	"github.com/SyamSolution/grule-service/internal/model"
	"github.com/SyamSolution/grule-service/internal/usecase"
	"github.com/SyamSolution/grule-service/internal/util"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type eligible struct {
	eligibleUsecase usecase.EligibleExecutor
	logger          config.Logger
}

type EligibleHandler interface {
	CheckEligibility(c *fiber.Ctx) error
}

func NewEligibleHandler(eligibleUsecase usecase.EligibleExecutor, logger config.Logger) EligibleHandler {
	return &eligible{
		eligibleUsecase: eligibleUsecase,
		logger:          logger,
	}
}

func (h *eligible) CheckEligibility(c *fiber.Ctx) error {
	var eligibleRequest model.Eligible
	if err := c.BodyParser(&eligibleRequest); err != nil {
		h.logger.Error("failed to parse request", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(model.ResponseWithoutData{
			Meta: model.Meta{
				Code:    400,
				Message: util.ERROR_NOT_FOUND_MSG,
			},
		})

	}

	isEligible, err := h.eligibleUsecase.CheckEligibility(&eligibleRequest)
	if err != nil {
		h.logger.Error("failed to check eligibility", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(model.ResponseWithoutData{
			Meta: model.Meta{
				Code:    500,
				Message: util.ERROR_BASE_MSG,
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Data: isEligible,
		Meta: model.Meta{
			Code:    200,
			Message: util.SUCCESS_RESPONSE_MSG,
		},
	})
}
