package apps

import (
	"github.com/hpazk/go-ticketing/apps/event"
	"github.com/hpazk/go-ticketing/apps/report"
	"github.com/hpazk/go-ticketing/apps/transaction"
	"github.com/hpazk/go-ticketing/apps/user"
	"github.com/hpazk/go-ticketing/helper"
	"github.com/hpazk/go-ticketing/routes"
	"github.com/labstack/echo/v4"
)

func AppInit(e *echo.Echo) {
	// Database
	// db := database.GetDbInstance()
	// dbMigration := database.GetMigrations(db)
	// err := dbMigration.Migrate()
	// if err != nil {
	// 	fmt.Println("migrations failed.", err)
	// } else {
	// 	fmt.Println("Migrations did run successfully")
	// }

	// Apps
	userApp := user.App{}
	eventApp := event.App{}
	tsxApp := transaction.App{}
	reportApp := report.App{}
	userApp.UseApp()
	eventApp.UseApp()
	tsxApp.UseApp()
	reportApp.UseApp()

	// Route
	handlers := []helper.Handler{
		&user.App{},
		&event.App{},
		&transaction.App{},
		&report.App{},
	}

	routes.DefineApiRoutes(e, handlers)
	// go scheduler.Scheduler()
}
