package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	ID        uint    `gorm:"primary_key" json:"id"`
	Username  string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Roles     []Role  `gorm:"many2many:user_roles;" json:"roles"`
	Groups    []Group `gorm:"many2many:user_groups;" json:"groups"`
	CreatedAt int64   `json:"created_at"`
	UpdatedAt int64   `json:"updated_at"`
	DeletedAt int64   `json:"deleted_at"`
}

type Role struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	DeletedAt   int64  `json:"deleted_at"`
}

type Group struct {
	ID            uint   `gorm:"primary_key" json:"id"`
	Name          string `json:"name"`
	ParentGroupID uint   `json:"parent_group_id"`
	ChildGroupIDs []uint `gorm:"-" json:"child_group_ids"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
	DeletedAt     int64  `json:"deleted_at"`
}

// Connexion à la DB

var db *gorm.DB

func connectDB() {
	var err error

	db, err = gorm.Open("postgres", "host=localhost user=myuser dbname=mydb sslmode=disable password=mypassword")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

}

// Fonctions pour auth JWT (signUp + Login)

func signUp(c *gin.Context) {

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
			"error": "Invalid email or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
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

func main() {
	connectDB()
	db.AutoMigrate(&User{})

	r := gin.Default()
	r.POST("/signup", signUp)
	r.POST("/login", login)

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", getUserList)
		userRoutes.GET("/:id", getUser)
		userRoutes.POST("/", createUser)
		userRoutes.PUT("/:id", updateUser)
		userRoutes.DELETE("/:id", deleteUser)
	}

	// Role endpoints
	roleRoutes := r.Group("/roles")
	{
		roleRoutes.GET("/", getRoleList)
		roleRoutes.GET("/:id", getRole)
		roleRoutes.POST("/", createRole)
		roleRoutes.PUT("/:id", updateRole)
		roleRoutes.DELETE("/:id", deleteRole)
	}

	// Group endpoints
	groupRoutes := r.Group("/groups")
	{
		groupRoutes.GET("/", getGroupList)
		groupRoutes.GET("/:id", getGroup)
		groupRoutes.POST("/", createGroup)
		groupRoutes.PUT("/:id", updateGroup)
		groupRoutes.DELETE("/:id", deleteGroup)
	}

	// Start the server
	r.Run(":8080")

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

//  Function endpoint user

// getUserList retourne la liste de tous les utilisateurs
func getUserList(c *gin.Context) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, users)
}

// getUser retourne un utilisateur par son ID
func getUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}

// createUser crée un nouvel utilisateur
func createUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := db.Create(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, user)
}

// updateUser met à jour un utilisateur existant
func updateUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := db.Save(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}

// deleteUser supprime un utilisateur existant
func deleteUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if err := db.Delete(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

//  Function endpoint role

// getRoleList retourne la liste de tous les rôles
func getRoleList(c *gin.Context) {
	var roles []Role
	if err := db.Find(&roles).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, roles)
}

// getRole retourne un rôle par son ID
func getRole(c *gin.Context) {
	var role Role
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&role).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, role)
}

// createRole crée un nouveau rôle
func createRole(c *gin.Context) {
	var role Role
	if err := c.BindJSON(&role); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := db.Create(&role).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, role)
}

// UpdateRole met à jour un rôle existant
func updateRole(c *gin.Context) {
	id := c.Param("id")
	var role Role
	if err := db.Where("id = ?", id).First(&role).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if err := c.BindJSON(&role); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	db.Save(&role)
	c.JSON(http.StatusOK, role)
}

// DeleteRole supprime un rôle par son ID
func deleteRole(c *gin.Context) {
	id := c.Param("id")
	var role Role
	if err := db.Where("id = ?", id).Delete(&role).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Le rôle a été supprimé avec succès"})
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

//  Function endpoint groupe.

// getGroupList retourne la liste de tous les groupes.
func getGroupList(c *gin.Context) {
	var groups []Group
	if err := db.Find(&groups).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to list groups"})
		return
	}
	c.JSON(http.StatusOK, groups)
}

// getRole retourne un groupe par son ID
func getGroup(c *gin.Context) {
	var group Group
	if err := db.First(&group, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "group not found"})
		return
	}
	c.JSON(http.StatusOK, group)
}

// createGroup crée un nouveau rôle
func createGroup(c *gin.Context) {
	var group Group
	if err := c.BindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided"})
		return
	}
	if err := db.Create(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create group"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": group})
}

// updateGroup met à jour un groupe existant
func updateGroup(c *gin.Context) {
	var group Group
	if err := db.First(&group, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "group not found"})
		return
	}
	if err := c.BindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group data"})
		return
	}
	if err := db.Save(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to update group"})
		return
	}
	c.JSON(http.StatusOK, group)
}

// deleteGroup supprime un groupe par son ID
func deleteGroup(c *gin.Context) {
	var group Group
	if err := db.First(&group, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "group not found"})
		return
	}
	if err := db.Delete(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to delete group"})
		return
	}
	c.Status(http.StatusNoContent)
}
