package middlewares

import (
	"net/http"

	"monopoly-Sandhu-Sahil/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized, Token Expired", err)
			c.Abort()
			return
		}
		c.Next()
	}
}

// var (
// 	RegixEmail              = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
// 	RegixAlphaWithSpecial   = regexp.MustCompile(`/^[a-zA-Z._^%$&#!~@, -]*$/`)
// 	RegixAlphabets          = regexp.MustCompile(`/^[a-zA-Z ]*$/`)
// 	RegixAlphabetsWithDots  = regexp.MustCompile(`/^[a-zA-Z. ]*$/`)
// 	RegixAlphabetsWithComma = regexp.MustCompile(`/^[a-zA-Z, ]*$/`)
// 	RegixAlphabetsWithSlash = regexp.MustCompile(`/^[a-zA-Z0-9/.& ]*$/`)
// 	RegixPassword           = regexp.MustCompile(`/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[|)(@\<{}>[\]/$!%*?:;.,=&_#~"'"^+-])[A-Za-z\d|)(@\<{}>[\]/$!%*?:;.,=&_#~"'"^+-]{8,}$/`)
// 	RegixPincode            = regexp.MustCompile(`/^[1-9][0-9]{5}$/`)
// 	RegixAlphanumeric       = regexp.MustCompile(`[^a-zA-Z0-9]+`)
// 	RegixPositiveInt        = regexp.MustCompile(`/^\d*$/`)
// 	RegixIFSC               = regexp.MustCompile(`/^([A-Za-z]{4}0[A-Za-z0-9]{6})$/`)
// )

// func ValidationUserLogin() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var varifyUser models.User
// 		if err := ctx.ShouldBindJSON(&varifyUser); err != nil {
// 			ctx.String(http.StatusUnauthorized, "Unauthorized, Required fields not given")
// 			ctx.Abort()
// 			return
// 		}
// 		fmt.Print(varifyUser)
// 		fmt.Print(RegixAlphabets.MatchString(varifyUser.Password))
// 		if !RegixAlphabets.MatchString(varifyUser.UserName) {
// 			ctx.String(http.StatusUnauthorized, "Unauthorized, user-name must be alphabets")
// 			ctx.Abort()
// 			return
// 		}
// 		ctx.Next()
// 	}
// }
