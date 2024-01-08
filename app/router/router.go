package router

import (
	"go-echo/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4" // echojwtという別名を与える
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New() // echoのインスタンスを作成

	// CORS対応
	// AllowOrigins: 許可するエンドポイント
	// AllowHeaders: Origin, Content-Type, Accept, Access-Control-Allow-Headers, X-CSRF-Token
	// AllowMethods: "GET", "PUT", "POST", "DELETE"
	// AlloCredentials: Cookieの送受信を可能にする
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	// CSRF対策
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",                     // CSRFのcookieパス
		CookieDomain:   os.Getenv("API_DOMAIN"), // CSRFのcookieドメイン
		CookieHTTPOnly: true,                    // CSRF cookieがhttpオンリーかどうか
		CookieSameSite: http.SameSiteNoneMode,   // こちらはsecure modeがtrueになるためPostmanで動作確認する際はコメントアウトする
		// CookieSameSite: http.SameSiteDefaultMode, // こちらはPostmanで動作確認する時に使用する
		// CookieMaxAge:   60, // トークンの有効期限
	}))

	// ユーザーAPI
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	e.GET("/csrf", uc.CsrfToken)

	// tasksグループを作成し、token認証していないとアクセスできないように設定
	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))

	// タスクAPI
	// :taskIdでURLパラメータの取得
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)

	return e
}
