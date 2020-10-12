package supervisor

import (
	"users_example/internal/supervisor/developer"
)

type Crud interface {
	Create(d *developer.Developer) Error
	Read(id string) (Crud, Error)
	Update(d *developer.Developer) Error
	Delete(id string) Error
}
