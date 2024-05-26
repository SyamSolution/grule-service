package usecase

import (
	"github.com/SyamSolution/grule-service/config"
	"github.com/SyamSolution/grule-service/internal/model"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/engine"
)

type discountUsecase struct {
	logger        config.Logger
	knowledgeBase *ast.KnowledgeBase
}

type DiscountExecutor interface {
	CheckDiscount(eligible *model.Discount) (float32, error)
}

func NewDiscountUsecase(logger config.Logger, knowledgeBase *ast.KnowledgeBase) DiscountExecutor {
	return &discountUsecase{
		logger:        logger,
		knowledgeBase: knowledgeBase,
	}
}

func (uc *discountUsecase) CheckDiscount(discount *model.Discount) (float32, error) {
	dataCtx := ast.NewDataContext()
	err := dataCtx.Add("D", discount)
	if err != nil {
		return 0, err
	}

	engine := engine.NewGruleEngine()
	err = engine.Execute(dataCtx, uc.knowledgeBase)
	if err != nil {
		return 0, err
	}

	return discount.DiscountAmount, nil
}
