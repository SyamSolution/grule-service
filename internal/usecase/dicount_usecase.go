package usecase

//type discountUsecase struct {
//	logger config.Logger
//}
//
//type DiscountExecutor interface {
//	CheckDiscount(eligible *model.Discount) (bool, error)
//}
//
//func NewDiscountUsecase(logger config.Logger) DiscountExecutor {
//	return &discountUsecase{logger: logger}
//}

//func (uc *discountUsecase) CheckDiscount(eligible *model.Discount) (bool, error) {
//	dataCtx := ast.NewDataContext()
//	err := dataCtx.Add("E", eligible)
//	if err != nil {
//		return false, err
//	}
//
//	knowledgeLibrary := ast.NewKnowledgeLibrary()
//	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)
//
//	fileRes := pkg.NewFileResource("./config/eligible.grl")
//	err = ruleBuilder.BuildRuleFromResource("EligibleRules", "0.0.1", fileRes)
//	if err != nil {
//		return false, err
//	}
//
//	knowledgeBase, _ := knowledgeLibrary.NewKnowledgeBaseInstance("EligibleRules", "0.0.1")
//
//	engine := engine.NewGruleEngine()
//	err = engine.Execute(dataCtx, knowledgeBase)
//	if err != nil {
//		return false, err
//	}
//
//	return , nil
//}
