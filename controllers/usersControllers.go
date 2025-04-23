package controllers

import (
	"fmt"
	"ic-project/initializers"
	"ic-project/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context){
	//Get the email/ password of req body
	var body struct{
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}
	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}
	//Create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}
	//Respond
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context){
	//Get the email and pass off req body
	var body struct{
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}
	//Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email or Password",
		})

		return
	}
	//Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email or Password",
		})

		return
	}
	/// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	//fmt.Printf("Token before signing: %+v\n", token) // Afișează structura
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET"))) 
	if err != nil {
		fmt.Println("JWT Error:", err) // Adaugă acest log
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	// send it back

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"Authorization",
		tokenString,
		3600 * 24 * 30, // Expirare (30 de zile)
		"/",            // Path
		"",             // Domain
		false,          // Secure (false pentru localhost)
		true,           // HttpOnly
	)
	c.JSON(http.StatusOK, gin.H{
	//"token" : tokenString,
	})

}

func Validate(c *gin.Context) {
    defer func() {
        if r := recover(); r != nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Nu ești autentificat",
            })
        }
    }()

    user, exists := c.Get("user")
    if !exists {
        panic("User not found in context") // Acest panic va fi recuperat
    }

    c.JSON(http.StatusOK, gin.H{
        "message": user,
    })
}