package controllers

import (
	"log"
	"net/http"
	"net/url"
	ent "prauth/entities"
	"prauth/models"
	"prauth/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	AppCtx *ent.AppCtx
	Dbs *services.DataBaseService
}

func (r AuthController) Auth(ctx *gin.Context){
	
	// session, _ := r.AppCtx.CookieStore.Get(ctx.Request, "prauth_session")
	// session.Values["user_id"] = http.StatusSeeOther
	// session.Save(ctx.Request, ctx.Writer)
	// fmt.Println(session.Values)
	ctx.HTML(http.StatusOK, "auth.tmpl", gin.H{
		"Alert": services.CreateAlert(ctx.Query("alert")),
	})
}

func (r AuthController) Signup(ctx *gin.Context) {
	var err error
	email := ctx.PostForm("email")
  
	u := models.User{
		Email: email,
		Username: email,
	}

	pwd1 := ctx.PostForm("password")
	pwd2 := ctx.PostForm("confirm-password")

	if (!services.IsValidEmail(email)){
		ctx.Redirect(http.StatusSeeOther, "/auth?alert=" + url.QueryEscape("e-Invalid User Email"))
		return
	}else if (!services.IsStrongPassword(pwd1)){
		ctx.Redirect(http.StatusSeeOther, "/auth?alert=" + url.QueryEscape("e-Weak User Password"))
		return
	}else if (pwd1 != pwd2){
		ctx.Redirect(http.StatusSeeOther, "/auth?alert=" + url.QueryEscape("e-User User Password Don't Match"))
		return
	}

	u.Password, err = services.HashPassword(pwd1)
	if err != nil{
		log.Println("Error Hashing New User Password: ",err)
	}

	if err := gorm.G[models.User](r.Dbs.DB).Create(r.Dbs.Ctx, &u); err != nil{
		log.Println("Error Creating User: ", err)
	}

	// Store user session
	session, _ := r.AppCtx.CookieStore.Get(ctx.Request, "prauth_session")
	session.Values["user_id"] = u.ID
	session.Save(ctx.Request, ctx.Writer)

	ctx.Redirect(http.StatusSeeOther, "/")
	
}

func (r AuthController) Signin(ctx *gin.Context){
	email := ctx.PostForm("email")
	pwd := ctx.PostForm("password")

	if !services.IsValidEmail(email){
		ctx.Redirect(http.StatusSeeOther, "/auth?authType=signin&alert=" + url.QueryEscape("e-Invalid User Email"))
	}else if (!services.IsStrongPassword(pwd)){
		ctx.Redirect(http.StatusSeeOther, "/auth?authType=signin&alert=" + url.QueryEscape("e-Invalid Password"))
		return
	}
  
	u, err := gorm.G[models.User](r.Dbs.DB).Where("email = ?", email).First(r.Dbs.Ctx)
	if err != nil{
		log.Println(err)
	}

	log.Println(u)

	if u.ID == 0 || !services.CheckPassword(u.Password, pwd){
		ctx.Redirect(http.StatusSeeOther, "/auth?authType=signin&alert=" + url.QueryEscape("e-Invalid UserIdentifier or Password"))
		return
	}

	// Store user session
	session, _ := r.AppCtx.CookieStore.Get(ctx.Request, "prauth_session")
	session.Values["user_id"] = u.ID
	session.Save(ctx.Request, ctx.Writer)

	ctx.Redirect(http.StatusSeeOther, "/")
}

func (r AuthController) Signout(ctx *gin.Context){

	session, _ := r.AppCtx.CookieStore.Get(ctx.Request, "prauth_session")
	session.Options.MaxAge = -1 // Mark for deletion
	session.Save(ctx.Request, ctx.Writer)
	ctx.Redirect(http.StatusPermanentRedirect, "/auth")
}

