package middleware

import  ent "prauth/entities"

type Middleware struct{
	AppCtx *ent.AppCtx
}


// func SetSession(c *gin.Context) {
// 	sess := c.MustGet("session").(*sessions.Session)
// 	sess.Values["user"] = "tristar"
// 	sess.Save(c.Request, c.Writer)
// 	c.JSON(http.StatusOK, gin.H{"message": "session set"})
// }
//
// func GetSession(c *gin.Context) {
// 	sess := c.MustGet("session").(*sessions.Session)
// 	user := sess.Values["user"]
// 	c.JSON(http.StatusOK, gin.H{"user": user})
// }
