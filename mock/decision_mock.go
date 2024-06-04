package mock

import (
	"github.com/gorules/zen-go"
	"github.com/stretchr/testify/mock"
)

type MockDecision struct {
	mock.Mock
}

func (m *MockDecision) Evaluate(context any) (*zen.EvaluationResponse, error) {
	args := m.Called(context)
	return args.Get(0).(*zen.EvaluationResponse), args.Error(1)
}

func (m *MockDecision) EvaluateWithOpts(context any, options zen.EvaluationOptions) (*zen.EvaluationResponse, error) {
	args := m.Called(context, options)
	return args.Get(0).(*zen.EvaluationResponse), args.Error(1)
}

func (m *MockDecision) Dispose() {
	m.Called()
}
