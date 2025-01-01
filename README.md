# Go Test Demo

This project includes code snippets and instructions to create a simple service and write rudimentary unittests for
it. Our goal is to have a working service for fetching items by id while mocking the repository layer in our tests.

## Running Tests
1. Install `mockery` following [mockery-installation](https://vektra.github.io/mockery/latest/installation/) guide. 
For example, on UNIX systems:
    ```
    $ brew install mockery
    ```
2. Install dependencies:
    ```
   $ go mod tidy
   ```
2. Generate mocks(mocks are already included in this repo) using the following command:
    ```
    $ mockery \
        --name=Repository \
        --dir=feature \
        --output=mocks \
        --with-expecter
    ```
   
   Explanation of flags:
   * `--name`: name of the interface to mock
   * `--dir`: location of the file containing the interface
   * `--output`: location of generated mocks
   * `--with-expecter`: augments mocks to include `EXPECT` method, and provide compile-time verification of method calls 
3. Run tests
    ```
    $ go test ./... -count=1
    ```
   
# Sample Test
Given repository
```go
type Repository interface {
	Get(models.ID) (*models.Feature, error)
}
```
and service
```go
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
```
We can mock the repository and write a simple test for the service as follows:
```go
func TestGetById_ReturnsExistingFeature(t *testing.T) {
	// given
	repo := mocks.NewRepository(t) // creating a new mock instance of Repository
	id := models.ID(1)
	repo.EXPECT().Get(id).Return(&models.Feature{}, nil) // specify expectations, expectations are automatically asserted
	                                                    // at the end of each test

	svc := ServiceImpl{Repo: repo}

	// when
	feature := svc.GetById(id)

	// then
	assert.NotNil(t, feature) // more assertions on the result
}
```