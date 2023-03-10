package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialisation de l'API

	r := gin.Default()

	r.GET("/users", usersPage)
	r.POST("/users", usersPost)
	r.PUT("/users", usersPut)
	r.DELETE("/users", usersDelete)

	r.GET("/roles", rolesPage)
	r.POST("/roles", rolesPost)
	r.PUT("/roles", rolesPut)
	r.DELETE("/roles", rolesDelete)

	r.GET("/groups", groupsPage)
	r.POST("/groups", groupsPost)
	r.PUT("/groups", groupsPut)
	r.DELETE("/groups", groupsDelete)

	r.POST("/auth", authJWT)

	r.Run()

}

// Fonctions pour le contenu des pages users

func usersPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Page des utilisateurs GET",
	})
}

func usersPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Page des utilisateurs POST",
	})
}

func usersPut(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Page des utilisateurs PUT",
	})
}

func usersDelete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Page des utilisateurs DELETE",
	})
}

// Fonctions pour le contenu des pages roles

func rolesPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Page des roles GET",
	})
}

func rolesPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Page des roles POST",
	})
}

func rolesPut(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Page des roles PUT",
	})
}

func rolesDelete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Page des roles DELETE",
	})
}

// Fonctions pour le contenu des pages groups

func groupsPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Page des groups GET",
	})
}

func groupsPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Page des groups POST",
	})
}

func groupsPut(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Page des groups PUT",
	})
}

func groupsDelete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Page des groups DELETE",
	})
}

// Fonction pour auth

func authJWT(c *gin.Context) {
	messageHTML := "AJOUT AUTHENTIFICATION"
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(messageHTML))
	return
}
