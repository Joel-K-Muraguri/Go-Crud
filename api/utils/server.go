package utils

import "github.com/Joel-K-Muraguri/Go-Crud/api/controller"

var server = controller.Server{}

func Run(){

	server.Initialize()

	server.Run(":8080")

}