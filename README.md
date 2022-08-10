# go fiber boilerplate

## Stack:
- go, fiber
- psql
- adminer
- docker-compose


### Routes
    * app.Post("/register", controllers.Register)
    * app.Post("/login", controllers.Login)
    * app.Get("/user", controllers.User)
    * app.Post("/logout", controllers.Logout)

### Models
    * User