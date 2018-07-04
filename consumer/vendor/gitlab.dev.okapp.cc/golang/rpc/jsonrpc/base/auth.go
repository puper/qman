package base

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthProvider interface {
	GetSecret(string) (string, error)
}

type AuthParams struct {
	AppKey       string `json:"app"`
	Sign         string `json:"sign"`
	Time         int64  `json:"ts"`
	RandomString string `json:"random"`
}

func AuthMiddleware(provider AuthProvider, validDuration float64) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			err        error
			authFailed = NewResponse(nil, AuthFailed, nil)
		)
		authv, err := base64.StdEncoding.DecodeString(ctx.GetHeader("X-JSONRPC-AUTH"))
		if err != nil {
			ctx.JSON(200, authFailed)
			return
		}
		auth := new(AuthParams)
		err = json.Unmarshal([]byte(authv), auth)
		if err != nil {
			ctx.JSON(200, authFailed)
			return
		}
		if math.Abs(float64(time.Now().Unix()-auth.Time)) > validDuration {
			ctx.JSON(200, authFailed)
			return
		}
		secret, err := provider.GetSecret(auth.AppKey)
		if err != nil {
			ctx.JSON(200, authFailed)
			return
		}
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(200, authFailed)
			return
		}
		ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		sign := SignData(auth.AppKey, strconv.Itoa(int(auth.Time)), auth.RandomString, secret, string(body))
		if sign != auth.Sign {
			ctx.JSON(200, authFailed)
			return
		}
		ctx.Set("jsonrpc.appKey", auth.AppKey)
		ctx.Next()
	}
}

func SignData(params ...string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(strings.Join(params, "")))
	cipherStr := md5Ctx.Sum(nil)
	result := hex.EncodeToString(cipherStr)
	return result[0:16]
}
