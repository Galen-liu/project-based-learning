package url_shortener

import (
		"github.com/gin-gonic/gin"
		"net/http"
		"net/url"
		"crypto/md5"
		"encoding/base64"
		"github.com/Galen-liu/project-based-learning/go/url_shortener/services/redis_service"
)

type OriginUrl struct {
				Url string `body:"url" binding:"required,url"`
}

type ShortenedUrl struct {
				Id string `uri:"id" binding:"required,base64url"`
}

func CreateShortenedUrl(ctx *gin.Context) {
	var originUrl OriginUrl
	if ctx.BindJSON(&originUrl) != nil {
		return
	}

	encodedUrl := url.PathEscape(originUrl.Url)

	hash := md5.Sum([]byte(encodedUrl))
	id := base64.URLEncoding.EncodeToString(hash[:])

	redisErr := redis_service.AddShortenUrlMap(id, encodedUrl)
	if redisErr != nil {
		ctx.AbortWithError(http.StatusInternalServerError, redisErr)
		return
	}


	ctx.IndentedJSON(http.StatusOK, gin.H{
		"originUrl": originUrl.Url,
		"shortenedUrl": "http://localhost:8080/shorten-url/" + id,
	})
}


func Redirect2RealUrl(ctx *gin.Context) {
	var shortenedUrl ShortenedUrl
	if err:=ctx.ShouldBindUri(&shortenedUrl); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if len(shortenedUrl.Id) == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	originUrl, redisErr := redis_service.GetShortenUrlMap(shortenedUrl.Id)
	if redisErr != nil {
		ctx.AbortWithError(http.StatusInternalServerError, redisErr)
		return
	}

	if len(originUrl) == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "NotFound",
		})
		return
	}

	decodedUrl, _ := url.PathUnescape(originUrl)
	ctx.Redirect(http.StatusFound, decodedUrl)
}

