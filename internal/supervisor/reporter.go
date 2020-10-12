package supervisor

import (
	"users_example/internal/supervisor/developer"
)

type LazyReporter interface {
	LazyReport() (developer.Report, Error)
}
