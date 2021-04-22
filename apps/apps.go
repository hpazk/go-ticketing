package apps

import (
	"github.com/hpazk/go-booklib/apps/event"
	"github.com/hpazk/go-booklib/apps/transaction"
	"github.com/hpazk/go-booklib/apps/user"
	"github.com/hpazk/go-booklib/database"
	"github.com/hpazk/go-booklib/helper"
	"github.com/hpazk/go-booklib/routes"
	"github.com/labstack/echo/v4"
)

func AppInit(e *echo.Echo) {
	// Database
	db := database.GetDbInstance()
	// err := db.AutoMigrate(
	// 	&user.User{},
	// 	&event.Event{},
	// 	&transaction.Transaction{},
	// )
	// // dbMigration := database.GetMigrations(db)
	// // err := dbMigration.Migrate()
	// if err == nil {
	// 	fmt.Println("Migrations did run successfully")
	// } else {
	// 	fmt.Println("migrations failed.", err)
	// }

	// Apps
	userApp := user.Init(db)
	eventApp := event.Init(db)
	tsxApp := transaction.Init(db)
	userApp.UseApp()
	eventApp.UseApp()
	tsxApp.UseApp()

	// Route
	handlers := []helper.Handler{
		&user.App{},
		&event.App{},
		&transaction.App{},
	}

	routes.DefineApiRoutes(e, handlers)
}
