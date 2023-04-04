package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
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
	var err error
	db, err = gorm.Open("postgres", connStr)

	if err != nil {
		panic(fmt.Sprintf("failed to connect to the database: %v", err))
	}
	defer db.Close()

	// Set up Gin router
	router := gin.Default()

	// user endpoints
	users := router.Group("/users")
	{
		users.Use(requireAuth)
		users.GET("/", getUsersList(db))
		users.GET("/:id", getUser(db))
		users.POST("/", createUser(db))
		users.PUT("/:id", updateUser(db))
		users.DELETE("/:id", deleteUser(db))
	}

	// Role endpoints
	roles := router.Group("/roles")
	{
		roles.Use(requireAuth)
		roles.GET("/", getRolesList(db))
		roles.GET("/:id", getRole(db))
		roles.POST("/", createRole(db))
		roles.PUT("/:id", updateRole(db))
		roles.DELETE("/:id", deleteRole(db))
	}

	// Group endpoints
	groups := router.Group("/groups")
	{
		groups.Use(requireAuth)
		groups.GET("/", getGroupsList(db))
		groups.GET("/:id", getGroup(db))
		groups.POST("/", createGroup(db))
		groups.PUT("/:id", updateGroup(db))
		groups.DELETE("/:id", deleteGroup(db))
	}

	// Auth endpoints
	router.POST("/signup", signup)
	router.POST("/login", login)

	// Start the server
	router.Run(":8080")

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

//  Function endpoint user

// getUsersList donne la liste des utilisateurs
func getUsersList(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

// getUser fetches a single user by their ID
func getUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user User
		if err := db.Where("id = ?", id).First(&user).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
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
func getRolesList(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var roles []Role
		if err := db.Find(&roles).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching roles"})
			return
		}
		c.JSON(http.StatusOK, roles)
	}
}

// getRole fetches a single role by their ID
func getRole(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var role Role
		if err := db.Where("id = ?", id).First(&role).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Role not found"})
			return
		}
		c.JSON(http.StatusOK, role)
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
func getGroupsList(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var groups []Group
		if err := db.Find(&groups).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching groups"})
			return
		}
		c.JSON(http.StatusOK, groups)
	}
}

// getGroup fetches a single group by their ID
func getGroup(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var group Group
		if err := db.Where("id = ?", id).First(&group).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Group not found"})
			return
		}
		c.JSON(http.StatusOK, group)
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

// Fonctions pour auth JWT (signup + Login)

func signup(c *gin.Context) {

	// get username/email/password du request body

	var body struct {
		Name     string
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

	user := User{Name: body.Name, Email: body.Email, Password: string(hash)}
	result := db.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Utilisateur enregistré, vous pouvez login",
	})
}

func login(c *gin.Context) {

	var body struct {
		Name     string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var user User
	db.First(&user, "name = ?", body.Name)

	if user.ID == 0 {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nom d'utilisateur ou mot de passe invalide",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nom d'utilisateur ou mot de passe invalide",
		})
		return
	}

	// generation du token JWT

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": user.ID,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
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
	c.JSON(http.StatusOK, gin.H{
		"message": "Vous êtes connecté",
	})

}

func requireAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	bearerToken := strings.Split(authHeader, " ")

	if len(bearerToken) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		return
	}

	tokenString := bearerToken[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("user_id", claims["id"])
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
}
