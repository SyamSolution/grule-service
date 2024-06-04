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

func TestCheckEligibility(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := mock_config.NewMockLogger(ctrl)
	mockDecision := new(mock_local.MockDecision)
	eligibleUsecase := NewEligibleUsecase(mockLogger, mockDecision)

	eligible := &model.Eligible{
		Day:   "1",
		Month: "1",
		Year:  "2022",
	}

	expectedResult := ResultData{Output: true}
	expectedResponse := zen.EvaluationResponse{
		Result: json.RawMessage(`{"output": true}`),
	}

	mockDecision.On("Evaluate", mock.Anything).Return(&expectedResponse, nil)

	result, err := eligibleUsecase.CheckEligibility(eligible)

	assert.NoError(t, err)
	assert.Equal(t, expectedResult.Output, result)

	mockDecision.AssertExpectations(t)
}
