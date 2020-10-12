package supervisor

import (
	"fmt"
	"github.com/google/uuid"
	"users_example/internal/supervisor/developer"
)

type Crud interface {
	Create(d *developer.Developer) error
	Read(id string) (*developer.Developer, error)
	Update(d *developer.Developer) error
	Delete(id string) error
}

type useCaseCRUD struct {
	repo developer.Repository
}

func (c useCaseCRUD) Create(d *developer.Developer) error {
	id := uuid.New()
	d.ID = id.String()

	if err := c.repo.Add(d); err != nil {
		return err
	}

	return nil
}

func (c useCaseCRUD) Read(id string) (*developer.Developer, error) {
	dev, err := c.repo.Get(id)
	if err != nil {
		return nil, err
	}

	if dev == nil {
		return nil, fmt.Errorf("developer:%s not found", id)
	}

	return dev, nil
}

func (c useCaseCRUD) Update(d *developer.Developer) error {
	panic("implement me")
}

func (c useCaseCRUD) Delete(id string) error {
	panic("implement me")
}
