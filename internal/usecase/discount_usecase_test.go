package usecase

import (
	"encoding/json"
	"github.com/SyamSolution/grule-service/internal/model"
	mock_local "github.com/SyamSolution/grule-service/mock"
	mock_config "github.com/SyamSolution/grule-service/mock/config"
	"github.com/gorules/zen-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestCheckDiscount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := mock_config.NewMockLogger(ctrl)
	mockDecision := &mock_local.MockDecision{}
	discountUsecase := NewDiscountUsecase(mockLogger, mockDecision)

	discount := &model.Discount{
		IsContinentSoldOut: true,
		IsContinentDiff:    true,
	}

	expectedResult := ResultDataDiscount{Output: 10.0}
	expectedResponse := zen.EvaluationResponse{
		Result: json.RawMessage(`{"output": 10.0}`),
	}

	mockDecision.On("Evaluate", mock.Anything).Return(&expectedResponse, nil)

	result, err := discountUsecase.CheckDiscount(discount)

	assert.NoError(t, err)
	assert.Equal(t, expectedResult.Output, result)
}
