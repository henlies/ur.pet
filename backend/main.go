package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henlies/project/controller"
	"github.com/henlies/project/entity"
)

func main() {
	entity.SetupDatabase()

	app := fiber.New()
	app.Use(CORSMiddleware())

	// - signin
	app.Post("/signin", controller.Signin)

	// - เอาไว้เช็คข้อมูล
	app.Get("/roles", controller.ListRoles)

	// - ป้องกันข้อมูล
	// api := app.Group("")
	// protected := api.Use(middlewares.Authorizes())

	// - Addrss
	app.Get("/addresses", controller.ListAddresses)
	app.Get("/address/user/:id", controller.GetAddressUID)
	app.Post("/address", controller.CreateAddress)
	app.Patch("/address", controller.UpdateAddress)

	// - Pet
	app.Get("/pets", controller.ListPets)
	app.Get("/pet/user/:id", controller.GetPetUID)
	app.Post("/pet", controller.CreatePet)
	app.Patch("/pet", controller.UpdatePet)

	// - User
	app.Patch("/userdetail", controller.UpdateUser)
	app.Patch("/userpass", controller.UpdatePasswordUser)

	app.Get("/user/:id", controller.GetUser)
	app.Get("/usersactive", controller.ListUsersActive)
	app.Get("/usersnonactive", controller.ListUsersNonactive)
	app.Delete("/serviceprovider/:email", controller.DeleteServiceProvider)
	app.Delete("/serviceuser/:email", controller.DeleteServiceUser)
	app.Delete("/userapprove/:id", controller.ApproveUser)
	app.Delete("/provideractive/:email", controller.ActiveServiceProvider)
	app.Delete("/useractive/:email", controller.ActiveServiceUser)
	app.Post("/usersigninuse", controller.CreateUserSigninUse)
	app.Post("/usersigninjob", controller.CreateUserSigninJob)

	// - Admin
	app.Get("/admins", controller.ListAdmins)
	app.Get("/admin/:id", controller.GetAdmins)
	app.Post("/admin", controller.CreateAdmin)
	app.Patch("/admindetail", controller.UpdateAdmin)
	app.Patch("/adminpass", controller.UpdatePasswordAdmin)
	app.Delete("/admin/:id", controller.DeleteAdmin)

	// - Post
	// app.Get("/post/:id", controller.GetPostbyId)
	// app.Get("/poststart", controller.ListPostStart)
	// app.Get("/poststart/:id", controller.ListPostStartId)
	// app.Get("/posttrack/:id", controller.ListPostIdTrack)
	// app.Get("/poststatus/:id", controller.ListPostIdStatus)
	// app.Post("/post", controller.CreatePost)
	// app.Patch("/post", controller.UpdatePost)
	// app.Patch("/postchk", controller.CheckPost)
	// app.Delete("/post/:id", controller.DeletePost)

	app.Listen(":8080")

}

func CORSMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Response().Header.Set("Access-Control-Allow-Origin", "*")
		c.Response().Header.Set("Access-Control-Allow-Credentials", "true")
		c.Response().Header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Response().Header.Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Method() == "OPTIONS" {
			return c.SendStatus(204)
		}

		return c.Next()
	}
}
