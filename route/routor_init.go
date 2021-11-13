package route

import (
	"github.com/gin-gonic/gin"
	"music_api_server/middleware"
	"net/http"
)

// RouterInit - /api/v1/path....
func RouterInit(r *gin.Engine){
	//Init for all available route
	apiV1 := r.Group("/api/v1")
	UseRoute(apiV1) //User route
	r.Use(middleware.JWTAuth()).GET("/test/jwt" ,func(context *gin.Context) {
		//name := context.Request.Header.Get("name")
		userName := context.MustGet("userName")
		//fmt.Println(email)
		context.JSON(http.StatusOK, gin.H{"message":userName})
	})

}

//func logger() gin.HandlerFunc{
//	return func(ctx* gin.Context){
//		t := time.Time{}
//		ctx.Set("test_field","jackson testing")
//
//		ctx.Next()
//
//		latency := time.Since(t)
//		fmt.Println(latency)
//		log.Println(ctx.Writer.Status())
//	}
//}