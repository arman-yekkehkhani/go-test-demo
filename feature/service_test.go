package feature

import (
	"errors"
	"github.com/arman-yekkehkhani/go-test-demo/mocks"
	"github.com/arman-yekkehkhani/go-test-demo/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetById_ReturnsExistingFeature(t *testing.T) {
	// given
	repo := mocks.NewRepository(t)
	id := models.ID(1)
	repo.EXPECT().Get(id).Return(&models.Feature{}, nil)

	svc := ServiceImpl{Repo: repo}

	// when
	feature := svc.GetById(id)

	// then
	assert.NotNil(t, feature)
}

func TestGetById_GivenNonExistingId_ReturnsNil(t *testing.T) {
	// given
	repo := mocks.NewRepository(t)
	id := models.ID(1)
	repo.EXPECT().Get(id).Return(nil, errors.New("feature not found"))

	svc := ServiceImpl{Repo: repo}

	// when
	feature := svc.GetById(id)

	// then
	assert.Nil(t, feature)
}
