package controllers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	ent "prauth/entities"
	"prauth/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AppCtx *ent.AppCtx
	Dbs *services.DataBaseService
}

func (r AuthController) Auth(ctx *gin.Context){
	
	// session, _ := r.AppCtx.CookieStore.Get(ctx.Request, "prauth_session")
	// session.Values["user_id"] = 303
	// session.Save(ctx.Request, ctx.Writer)
	// fmt.Println(session.Values)
	fmt.Println("***",ctx.Query("msg"))
	ctx.HTML(http.StatusOK, "auth.tmpl", gin.H{
		"Msg": "working",
	})
}

func (r AuthController) Signup(ctx *gin.Context) {
	var err error

	u := ent.User{
		Dbs: r.Dbs,
		Email: ctx.PostForm("email"),
	}

	pwd1 := ctx.PostForm("password")
	pwd2 := ctx.PostForm("confirm-password")

	if (services.IsValidEmail(u.Email)){
		ctx.Redirect(303, "/auth?msg=" + url.QueryEscape("e Invalid User Email"))
		return
	}else if (services.IsStrongPassword(pwd1)){
		ctx.Redirect(303, "/auth?msg=" + url.QueryEscape("e Weak User Password"))
		return
	}else if (pwd1 != pwd2){
		ctx.Redirect(303, "/auth?msg=" + url.QueryEscape("e User User Password Don't Match"))
		return
	}

	u.PwdHash, err = services.HashPassword(pwd1)
	if err != nil{
		log.Println("Error Hashing New User Password: ",err)
	}

	if err := u.Create(); err != nil{
		log.Println("Error Creating User: ", err)
	}

	ctx.JSON(http.StatusOK,u)
	
}

func (r AuthController) Signin(ctx *gin.Context){
	u := ent.User{
		Dbs: r.Dbs,
		Email: ctx.PostForm("email"),
	}

	err := u.GetByEmail()
	if err != nil{
		log.Println("Error Getting User: ",err)
	}
	
	log.Println("***", u)
	if u.ID == 0 || services.CheckPassword(u.PwdHash, ctx.PostForm("password")){
		ctx.JSON(http.StatusOK, "Invalid User Name or Password")
	}

	ctx.JSON(http.StatusOK,u)
}

func (r AuthController) Signout(ctx *gin.Context){

	session, _ := r.AppCtx.CookieStore.Get(ctx.Request, "prauth_session")
	session.Options.MaxAge = -1 // Mark for deletion
	session.Save(ctx.Request, ctx.Writer)
	ctx.Redirect(http.StatusPermanentRedirect, "/auth")
}

