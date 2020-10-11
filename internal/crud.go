package internal

type CrudDeveloper interface {
	Create(d *Developer) DevApiError
	Read(id string) (CrudDeveloper, DevApiError)
	Update(d *Developer) DevApiError
	Delete(id string) DevApiError
}
