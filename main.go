package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func load() {
	init.connectDB()
}

func main() {

	r := gin.Default()

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

	// r.POST("/auth", authenticateUser)

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
