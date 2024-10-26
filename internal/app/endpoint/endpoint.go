package endpoint

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type DBService interface {
	GetAllUsers() ([]string, error)
	CheckUser(string, string) (bool, error)
}

type Endpoint struct {
	db DBService
}

func New(db DBService) *Endpoint {
	return &Endpoint{db: db}
}

func (ep *Endpoint) GetUsers(ctx *gin.Context) {
	users, _ := ep.db.GetAllUsers()
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (ep *Endpoint) Auth(ctx *gin.Context) {
	var authRequest AuthRequest
	if err := ctx.ShouldBindJSON(&authRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	// Извлечение значений
	login := authRequest.Login
	password := authRequest.Password
	/// Проверить в бд
	//ChekUser вернет id пользователя и установит в сессию
	if isExist, _ := ep.db.CheckUser(login, password); !isExist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"user not found": ""})
		ctx.Abort()
		return
	}
	// Генерация MD5 ключа для сессии
	hash := md5.Sum([]byte(password))
	md5Key := hex.EncodeToString(hash[:])
	// Создание сессии
	session := sessions.Default(ctx)
	session.Set("login", login)
	session.Set("md5Key", md5Key)
	session.Save()
	ctx.JSON(http.StatusBadRequest, gin.H{"ses": session.Get("md5Key")})

}
