package internal

type LazyReporter interface {
	LazyReport() DevApiError
}
