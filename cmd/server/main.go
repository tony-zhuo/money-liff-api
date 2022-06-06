package main

import "github.com/ZhuoYIZIA/money-liff-api/routes"

func main() {
	engine := routes.InitRoutes()
	panic(engine.Run())
}
