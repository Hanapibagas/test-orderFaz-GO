package service

import (
	"orderfaz-1/features/logistic"
	"orderfaz-1/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	t.Run("with result", func(t *testing.T) {
		repo := new(mocks.LogisticData)

		srv := NewLogistic(repo)
		inputData := []logistic.CoreLogistic{
			{
				LogisticName:    "JNE",
				Amount:          "10000",
				DastinationName: "JAKARTA",
				OriginName:      "BANDUNG",
				Duration:        "2-4",
			},
		}

		repo.On("SerachByQuery", "origin_name", "dastination_name").Return(inputData, nil).Once()

		resulr, err := srv.SerachByQuery("origin_name", "dastination_name")

		assert.NoError(t, err)
		assert.Equal(t, inputData, resulr)
		repo.AssertExpectations(t)
	})

	t.Run("Without Results", func(t *testing.T) {
		repo := new(mocks.LogisticData)

		srv := NewLogistic(repo)
		repo.On("SerachByQuery", "nonexistent", "nonexistent").Return([]logistic.CoreLogistic{}, nil).Once()

		results, err := srv.SerachByQuery("nonexistent", "nonexistent")

		assert.NoError(t, err)
		assert.Empty(t, results)
		repo.AssertExpectations(t)
	})
}
