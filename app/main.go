package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         uint        `gorm:"primary_key" json:"id"`
	Name       string      `json:"name"`
	Email      string      `gorm:"unique" json:"email"`
	Password   string      `json:"-"`
	Roles      []Role      `gorm:"many2many:user_roles;" json:"roles"`
	Groups     []Group     `gorm:"many2many:user_groups;" json:"groups"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	DeletedAt  *time.Time  `json:"deleted_at"`
	AuthTokens []AuthToken `json:"auth_tokens"`
}

type AuthToken struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	UserID    uint      `json:"-"`
}

type RefreshToken struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	UserID    uint      `json:"-"`
}

type Role struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type Group struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	Name          string     `json:"name"`
	ParentGroupID *uint      `json:"parent_group_id"`
	ChildGroupIDs []uint     `gorm:"-" json:"child_group_ids"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

var db *gorm.DB

func main() {

	// Load the .env file
	err1 := godotenv.Load()
	if err1 != nil {
		panic("Error loading .env file")
	}

	// Environment variables
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	// Connect to the PostgreSQL database
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbname)
	db, err := gorm.Open("postgres", connStr)

	if err != nil {
		panic(fmt.Sprintf("failed to connect to the database: %v", err))
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{}, &AuthToken{}, &RefreshToken{}, &Role{}, &Group{})

	// Set up Gin router
	router := gin.Default()

	// user endpoints
	users := router.Group("/users")
	{
		users.GET("/", getUsers(db))
		users.POST("/", createUser(db))
		users.PUT("/:id", updateUser(db))
		users.DELETE("/:id", deleteUser(db))
	}

	// Role endpoints
	roles := router.Group("/roles")
	{
		roles.GET("/", getRoles(db))
		roles.POST("/", createRole(db))
		roles.PUT("/:id", updateRole(db))
		roles.DELETE("/:id", deleteRole(db))
	}

	// Group endpoints
	groups := router.Group("/groups")
	{
		groups.GET("/", getGroups(db))
		groups.POST("/", createGroup(db))
		groups.PUT("/:id", updateGroup(db))
		groups.DELETE("/:id", deleteGroup(db))
	}

	// Start the server
	router.Run(":8080")

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

//  Function endpoint user

// User handlers
func getUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

// createUser crée un nouvel utilisateur
func createUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
			return
		}

		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
			return
		}
		c.JSON(http.StatusCreated, user)
	}
}

// updateUser met à jour un utilisateur existant
func updateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user User
		if err := db.Where("id = ?", id).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
			return
		}

		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// deleteUser supprime un utilisateur existant
func deleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user User
		if err := db.Where("id = ?", id).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		if err := db.Delete(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

//  Function endpoint role

// getRoles retourne la liste de tous les rôles et retourne un rôle par son ID
func getRoles(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var roles []Role
		if err := db.Find(&roles).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching roles"})
			return
		}
		c.JSON(http.StatusOK, roles)
	}
}

// createRole crée un nouveau rôle
func createRole(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var role Role
		if err := c.BindJSON(&role); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role data"})
			return
		}

		if err := db.Create(&role).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating role"})
			return
		}
		c.JSON(http.StatusCreated, role)
	}
}

// UpdateRole met à jour un rôle existant
func updateRole(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var role Role
		if err := db.Where("id = ?", id).First(&role).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
			return
		}

		if err := c.BindJSON(&role); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role data"})
			return
		}

		if err := db.Save(&role).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating role"})
			return
		}
		c.JSON(http.StatusOK, role)
	}
}

// DeleteRole supprime un rôle par son ID
func deleteRole(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var role Role
		if err := db.Where("id = ?", id).First(&role).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
			return
		}

		if err := db.Delete(&role).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting role"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Role deleted"})
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

//  Function endpoint groupe.

// getGroups retourne la liste de tous les groupes et un groupe par son ID
func getGroups(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var groups []Group
		if err := db.Find(&groups).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching groups"})
			return
		}
		c.JSON(http.StatusOK, groups)
	}
}

// createGroup crée un nouveau rôle
func createGroup(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var group Group
		if err := c.BindJSON(&group); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group data"})
			return
		}

		if err := db.Create(&group).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating group"})
			return
		}
		c.JSON(http.StatusCreated, group)
	}
}

// updateGroup met à jour un groupe existant
func updateGroup(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var group Group
		if err := db.Where("id = ?", id).First(&group).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
			return
		}

		if err := c.BindJSON(&group); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group data"})
			return
		}

		if err := db.Save(&group).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating group"})
			return
		}
		c.JSON(http.StatusOK, group)
	}
}

// deleteGroup supprime un groupe par son ID
func deleteGroup(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var group Group
		if err := db.Where("id = ?", id).First(&group).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
			return
		}

		if err := db.Delete(&group).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting group"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Group deleted"})
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Fonctions pour auth JWT (signUp + Login)

func signup(c *gin.Context) {

	// get email/password du request body

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// hash password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// creation user

	user := User{Email: body.Email, Password: string(hash)}
	result := db.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func login(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var user User
	db.First(&user, "email = ?", body.Email)

	if user.ID == 0 {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email ou mot de passe invalide",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email ou mot de passe invalide",
		})
		return
	}

	// generation du token JWT

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// signature et recup du token chiffré en string utilisant la var SECRET

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// on retourne le token (en cookie)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})

}

// message de confirmation de connexion
func validation(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message":  "Vous êtes connecté",
		"username": user,
	})

}

func requireAuth(c *gin.Context) {

	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])

		}

		return []byte(os.Getenv("SECRET")), nil

	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user User
		db.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)

		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
