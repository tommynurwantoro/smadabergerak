package domain

type Handler interface {
	General() GeneralHandler
	SmadaBergerak() SmadaBergerakHandler
}
