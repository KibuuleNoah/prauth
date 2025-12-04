package routes

import (
	"embed"
	"html/template"
	// "log"
	"prauth/controllers"
	"prauth/database"
	ent "prauth/entities"
	"prauth/middleware"
	"prauth/services"
	// "time"

	// ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	// "go.uber.org/zap"
)

func RunServer(templateFS *embed.FS, staticFS embed.FS, appCtx *ent.AppCtx) {
	router := gin.Default()

	// logger, err := zap.NewProduction()
	// if err != nil{
	// 	log.Fatal("Logger Error: ",err)
	// }

	// router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

  // Logs 
	router.SetFuncMap(services.TemplatesService{}.All())

	// Load templates from embedded FS
	tmpl, _ := template.ParseFS(templateFS, "templates/**/*.tmpl")
	router.SetHTMLTemplate(tmpl)

	// Serve embedded static files
	services.CustomServeStaticFS(router, staticFS)

	db, ctx := database.InitDB()

	//DATABASE SERVICE
	dbservice := services.DataBaseService{DB: db, Ctx: ctx}

	mw := middleware.Middleware{AppCtx: appCtx}
	// client routes
	clientCtrl := controllers.ClientController{AppCtx: appCtx}

	client := router.Group("/").Use(mw.AuthRequired())
	client.GET("", clientCtrl.Index)

	// api endpoints
	// apiCtrl := controllers.ApiController{DBService:dbservice}


	// AUTH ROUTES
	authCtrl := controllers.AuthController{AppCtx: appCtx, Dbs: &dbservice}
	auth := router.Group("/auth")
	{
		auth.GET("/",authCtrl.Auth)
		auth.POST("/signup", authCtrl.Signup)
		auth.POST("/signin", authCtrl.Signin)
		auth.GET("/signout", authCtrl.Signout)
	}
	
	
	// PROTECTED ROUTES
	// admin := r.Group("/admin")
	// a.Use(middleware.AuthRequired())
	// {
	// 	auth.GET("/dashboard", func(c *gin.Context) {
	// 		c.JSON(200, gin.H{"message": "Welcome admin"})
	// 	})
	// }

	// api endpoints
	// apiCtrl := controllers.ApiController{DBService:dbservice}

	// api version one
	// apiV1 := router.Group("/api/v1")
	// {
	// 	apiV1.GET("/schools/exam/ple/", apiCtrl.SchoolsPLE)
	// 	apiV1.GET("/schools/exam/uce/", apiCtrl.SchoolsUCE)
	// 	apiV1.GET("/schools/exam/uace/", apiCtrl.SchoolsUACE)
	// 	apiV1.GET("/school/samplesUCE/", apiCtrl.SampleUCE)
	// 	apiV1.Any("/testers", apiCtrl.Testers)
	// }

	router.Run(":" + services.GetServerPort())
}
