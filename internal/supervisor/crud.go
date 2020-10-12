package supervisor

import (
	"users_example/internal/supervisor/developer"
)

type Crud interface {
	Create(d *developer.Developer) error
	Read(id string) (Crud, error)
	Update(d *developer.Developer) error
	Delete(id string) error
}
