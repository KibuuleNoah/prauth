package controllers

import (
	"net/http"
	ent "prauth/entities"

	"github.com/gin-gonic/gin"
)

type ClientController struct {
	AppCtx *ent.AppCtx
}

func (r ClientController) Index(ctx *gin.Context){
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"msg": "working",
	})
}

