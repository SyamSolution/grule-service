package usecase

import (
	"encoding/json"

	"github.com/SyamSolution/grule-service/config"
	"github.com/SyamSolution/grule-service/internal/model"
	"github.com/gorules/zen-go"
)

type eligibleUsecase struct {
	logger   config.Logger
	ruleBase zen.Decision
}
type ResultData struct {
	Output bool `json:"output"`
}

type EligibleExecutor interface {
	CheckEligibility(eligible *model.Eligible) (bool, error)
}

func NewEligibleUsecase(logger config.Logger, rulebase zen.Decision) EligibleExecutor {
	return &eligibleUsecase{
		logger:   logger,
		ruleBase: rulebase,
	}
}

func (uc *eligibleUsecase) CheckEligibility(eligible *model.Eligible) (bool, error) {
	response, err := uc.ruleBase.Evaluate(map[string]any{"day": eligible.Day, "month": eligible.Month, "year": eligible.Year})
	if err != nil {
		panic(err)
	}
	var data ResultData
	if err := json.Unmarshal([]byte(response.Result), &data); err != nil {
		return false, err
	}
	return data.Output, nil
}
