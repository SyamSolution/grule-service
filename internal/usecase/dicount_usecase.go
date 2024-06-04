package usecase

import (
	"encoding/json"

	"github.com/SyamSolution/grule-service/config"
	"github.com/SyamSolution/grule-service/internal/model"
	"github.com/gorules/zen-go"
)

type discountUsecase struct {
	logger   config.Logger
	ruleBase zen.Decision
}

type ResultDataDiscount struct {
	Output float32 `json:"output"`
}

type DiscountExecutor interface {
	CheckDiscount(eligible *model.Discount) (float32, error)
}

func NewDiscountUsecase(logger config.Logger, rulebase zen.Decision) DiscountExecutor {
	return &discountUsecase{
		logger:   logger,
		ruleBase: rulebase,
	}
}

func (uc *discountUsecase) CheckDiscount(discount *model.Discount) (float32, error) {

	response, err := uc.ruleBase.Evaluate(map[string]any{"isContinentSoldOut": discount.IsContinentSoldOut, "isContinentDiff": discount.IsContinentDiff})
	if err != nil {
		panic(err)
	}
	var data ResultDataDiscount
	if err := json.Unmarshal([]byte(response.Result), &data); err != nil {
		return 0, err
	}
	return data.Output, nil
}
