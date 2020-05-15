package main

type Server interface {
	Initialize(configFile string)
	SetupAssets()
	SetupRoutes()
    ListenAndServe()
	SetupServer()
}

