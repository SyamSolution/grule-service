package handler

import (
	"bytes"
	"github.com/SyamSolution/grule-service/internal/model"
	"github.com/SyamSolution/grule-service/mock"
	mock_config "github.com/SyamSolution/grule-service/mock/config"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckDiscount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDiscountExecutor := mock.NewMockDiscountExecutor(ctrl)
	mockLogger := mock_config.NewMockLogger(ctrl)

	discountHandler := NewDiscountHandler(mockDiscountExecutor, mockLogger)

	discountRequest := &model.Discount{
		IsContinentSoldOut: true,
		IsContinentDiff:    true,
	}

	expectedDiscountAmount := float32(10.0)

	mockDiscountExecutor.EXPECT().CheckDiscount(discountRequest).Return(expectedDiscountAmount, nil)

	app := fiber.New()
	app.Post("/discount", discountHandler.CheckDiscount)

	resp, err := app.Test(httptest.NewRequest("POST", "/discount", bytes.NewBufferString(`{"IsContinentSoldOut":true,"IsContinentDiff":true}`)))

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(t, string(body), `"data":10`)
}
