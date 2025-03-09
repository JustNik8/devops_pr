package service

import (
	"context"
	"ecom/internal/domain"
	"ecom/internal/repository"
	"ecom/internal/repository/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllGoods(t *testing.T) {
	testGoods := []domain.Good{
		{ID: "test_id_1", Name: "test1", Price: 1000, Desc: "test good1", StockQuantity: 1, MeasureUnit: "kg"},
	}

	testCases := []struct {
		name          string
		filters       []domain.GormFilter
		ordersStr     string
		buildGoodRepo func() repository.GoodRepo
		expectedGoods []domain.Good
		isErrExpected bool
	}{
		{
			name:      "Get all goods without errors",
			filters:   []domain.GormFilter{},
			ordersStr: "",
			buildGoodRepo: func() repository.GoodRepo {
				mockRepo := mocks.NewGoodRepo(t)
				mockRepo.On("GetAllGoods", mock.Anything, mock.Anything, mock.Anything).
					Return(testGoods, nil).
					Once()

				return mockRepo
			},
			expectedGoods: testGoods,
			isErrExpected: false,
		},
		{
			name:      "Get all goods with error",
			filters:   []domain.GormFilter{},
			ordersStr: "",
			buildGoodRepo: func() repository.GoodRepo {
				mockRepo := mocks.NewGoodRepo(t)
				mockRepo.On("GetAllGoods", mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errors.New("error while getting all goods")).
					Once()

				return mockRepo
			},
			expectedGoods: nil,
			isErrExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			goodService := NewGoodService(
				tc.buildGoodRepo(),
			)

			actualGoods, err := goodService.GetAllGoods(context.Background(), tc.filters, tc.ordersStr)
			assert.Equal(t, tc.expectedGoods, actualGoods)
			assert.Equal(t, tc.isErrExpected, err != nil)
		})
	}

}
