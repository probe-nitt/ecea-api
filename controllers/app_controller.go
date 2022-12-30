package controllers

type AppController struct {
	User interface{ UserController }
}
