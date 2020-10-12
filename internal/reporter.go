package internal

type LazyReporter interface {
	LazyReport() (Report, DevApiError)
}
