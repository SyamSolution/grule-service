package usecase

import (
	"github.com/SyamSolution/grule-service/config"
	"github.com/SyamSolution/grule-service/internal/model"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/engine"
)

type eligibleUsecase struct {
	logger        config.Logger
	knowledgeBase *ast.KnowledgeBase
}

type EligibleExecutor interface {
	CheckEligibility(eligible *model.Eligible) (bool, error)
}

func NewEligibleUsecase(logger config.Logger, knowledgeBase *ast.KnowledgeBase) EligibleExecutor {
	return &eligibleUsecase{
		logger:        logger,
		knowledgeBase: knowledgeBase,
	}
}

func (uc *eligibleUsecase) CheckEligibility(eligible *model.Eligible) (bool, error) {
	dataCtx := ast.NewDataContext()
	err := dataCtx.Add("E", eligible)
	if err != nil {
		return false, err
	}

	engine := engine.NewGruleEngine()
	err = engine.Execute(dataCtx, uc.knowledgeBase)
	if err != nil {
		return false, err
	}

	return eligible.Eligibility, nil
}
