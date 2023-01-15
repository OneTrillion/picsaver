package util

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/gin-gonic/gin"
)

func GenerateOathCookie(ctx *gin.Context) string {

    cookieSize := make([]byte, 16)

    // Generate random bytes
    rand.Read(cookieSize)
    state := base64.URLEncoding.EncodeToString(cookieSize)

    // Fix localhost input (more general)
    ctx.SetCookie("oauthCookie", state, 3600, "/", "localhost", false, true)

    return state
}
