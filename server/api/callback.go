package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"picsaver/server/config"

	"github.com/gin-gonic/gin"
)

func (server *Server) GoogleCallback(ctx *gin.Context)  {
    oauthState, _ := ctx.Cookie("oauthstate")
    state := ctx.Request.FormValue("state")
    code := ctx.Request.FormValue("code")
    ctx.Request.Header.Add("content-type", "application/json")

    if state != oauthState {
        ctx.Redirect(http.StatusTemporaryRedirect, "/")
        log.Println("invalid oauth google state")
		return
    }

    token, err := config.AppConfig.GoogleLoginConfig.Exchange(
        context.Background(), code,
    )

    if err != nil {
        log.Printf("falied code exchange: %s", err.Error())
        return
    }

    res, exists := ctx.Get(config.OauthGoogleUrlAPI + token.AccessToken)
    
    if exists == false {
        log.Printf("failed getting user info: %s", err.Error())
		return 
    }

    //TODO: ta bort
    fmt.Println(res)

    //TODO: Jämför med databasen??
    //kanske lägg in i databasen??

    //TODO: skicka shiten till react frontend
    ctx.JSON(http.StatusOK, "Hello")

    /*
    TODO: ta bort sen 
    REFRENCE: 

    defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(rw, "failed read response: %s", err.Error())
		return
	}

    //TODO: skicka shiten till react frontend
	fmt.Fprintln(rw, string(contents))
    */
}
