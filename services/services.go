package services

import (
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoadJsonFile[T any](filePath string) T {
	file, err := os.ReadFile(filePath)

	var v T

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(file, &v); err != nil {
		log.Fatal(err)
	}

	return v
}

func GetServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func GetDBURL() string {
	url := os.Getenv("DB_URL")
	if url == "" {
		log.Fatal("Failed to get DB_URL Enviromental variable")
	}
	return url
}

func CustomServeStaticFS(r *gin.Engine, fsys fs.FS){
	fs.WalkDir(fsys, "static", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			r.GET("/" + path, func(c *gin.Context) {
				c.FileFromFS(path, http.FS(fsys))
			})
		}
		return nil
	})
}


func IsValidEmail(email string) bool {
    regex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
    return regex.MatchString(email)
}


func IsStrongPassword(plaintext string) bool {
    var (
        hasMinLen  = false
        hasUpper   = false
        hasLower   = false
        hasNumber  = false
        hasSpecial = false
    )
    s := []rune(plaintext)
    if len(s) >= 7 {
        hasMinLen = true
    }
    for _, char := range s {
        switch {
        case unicode.IsUpper(char):
            hasUpper = true
        case unicode.IsLower(char):
            hasLower = true
        case unicode.IsNumber(char):
            hasNumber = true
        case unicode.IsPunct(char) || unicode.IsSymbol(char):
            hasSpecial = true
        }
    }
    return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

// HashPassword hashes a plaintext password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// CheckPassword compares a bcrypt hashed password with a plaintext one.
func CheckPassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func CreateAlert(alertStr string) (alert map[string]string) {
	split := strings.Split(alertStr, "-")
	if len(split) < 2{
		return alert
	}
	alertTypes := map[string]string{"e":"error", "s": "success", "i": "info"}
	return map[string]string{"type": alertTypes[split[0]], "text": split[1]}
}

func LoginUser(){}
func LogoutUser(){}
