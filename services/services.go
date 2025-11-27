package services

import (
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"os"
	"regexp"

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


func IsStrongPassword(password string) bool {
    // One regex to check everything
    regex := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[^a-zA-Z0-9]).{8,}$`)
    return regex.MatchString(password)
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

func LoginUser(){}
func LogoutUser(){}
