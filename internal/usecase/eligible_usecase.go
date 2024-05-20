package usecase

import (
	"github.com/SyamSolution/grule-service/config"
	"github.com/SyamSolution/grule-service/internal/model"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type eligibleUsecase struct {
	logger config.Logger
}

type EligibleExecutor interface {
	CheckEligibility(eligible *model.Eligible) (bool, error)
}

func NewEligibleUsecase(logger config.Logger) EligibleExecutor {
	return &eligibleUsecase{logger: logger}
}

func (uc *eligibleUsecase) CheckEligibility(eligible *model.Eligible) (bool, error) {
	dataCtx := ast.NewDataContext()
	err := dataCtx.Add("E", eligible)
	if err != nil {
		return false, err
	}

	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

	fileRes := pkg.NewFileResource("./config/rule.grl")
	err = ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", fileRes)
	if err != nil {
		return false, err
	}

	knowledgeBase, _ := knowledgeLibrary.NewKnowledgeBaseInstance("TutorialRules", "0.0.1")

	engine := engine.NewGruleEngine()
	err = engine.Execute(dataCtx, knowledgeBase)
	if err != nil {
		return false, err
	}

	return eligible.Eligibility, nil
}
