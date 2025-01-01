package feature

import "github.com/arman-yekkehkhani/go-test-demo/models"

type Service interface {
	GetById(models.ID) *models.Feature
}

type ServiceImpl struct {
	Repo Repository
}

func (s *ServiceImpl) GetById(id models.ID) *models.Feature {
	f, err := s.Repo.Get(id)
	if err != nil {
		return nil
	}
	return f
}
