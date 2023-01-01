package controllers

type AppController struct {
	User interface{ UserController }
	Team interface{ TeamController }
}
