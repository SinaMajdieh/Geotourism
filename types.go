package main

type Server interface {
	SetupAssets()
	SetupRoutes()
    ListenAndServe()
	SetupServer()
}

