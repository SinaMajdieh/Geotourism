package internal

type Server interface {
	SetupAssets()
	SetupRoutes()
    ListenAndServe()
	SetupServer()
}

