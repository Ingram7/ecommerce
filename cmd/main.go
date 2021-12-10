package main

import (
	"ecommerce/app"
)

func main() {

	//ctx := logger.NewTagContext(context.Background(), "__main__")
	//ctx := context.Background()

	engine := app.NewEngine()
	engine.WarmUp()
	engine.Run()
}
