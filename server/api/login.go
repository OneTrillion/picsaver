package api

import (
	"net/http"
	"picsaver/server/config"
	"picsaver/server/util"

	"github.com/gin-gonic/gin"
)

func (server *Server) GoogleLogin(ctx *gin.Context)  {
    oauthCookieState := util.GenerateOathCookie(ctx) 

    url := config.AppConfig.GoogleLoginConfig.AuthCodeURL(oauthCookieState)
    ctx.Redirect(http.StatusTemporaryRedirect, url)
}
