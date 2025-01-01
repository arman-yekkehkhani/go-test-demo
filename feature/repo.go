package feature

import "github.com/arman-yekkehkhani/go-test-demo/models"

type Repository interface {
	Get(models.ID) (*models.Feature, error)
}
