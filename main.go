package main

import "github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/di"

func main() {
	app := di.Wire()
	app.Listen(":8080")
}
