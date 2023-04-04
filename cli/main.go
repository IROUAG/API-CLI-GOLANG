package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "Une CLI pour interagir avec l'API",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Use a subcommand to interact with the API. Run ' --help' for usage.")
		},
	}

	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Run the CLI as a server",
		Run:   serve,
	}
	rootCmd.AddCommand(serveCmd)

	// Login
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Connecter un utilisateur et recevoir un jeton d'authentification JWT et un jeton de rafraîchissement (refresh token)",
		Run:   login,
	}
	loginCmd.Flags().String("email", "", "L'adresse email de l'utilisateur")
	loginCmd.Flags().String("password", "", "Le mot de passe de l'utilisateur")
	rootCmd.AddCommand(loginCmd)

	// Refresh
	refreshCmd := &cobra.Command{
		Use:   "refresh",
		Short: "Renouveler un jeton d'authentification JWT à l'aide d'un jeton de rafraîchissement",
		Run:   refresh,
	}
	refreshCmd.Flags().String("refresh_token", "", "Le jeton de rafraîchissement")
	rootCmd.AddCommand(refreshCmd)

	// Logout
	logoutCmd := &cobra.Command{
		Use:   "logout",
		Short: "Supprimer un jeton d'authentification JWT et de rafraîchissement",
		Run:   logout,
	}
	logoutCmd.Flags().String("access_token", "", "Le jeton d'authentification à supprimer")
	logoutCmd.Flags().String("refresh_token", "", "Le jeton de rafraîchissement à supprimer")
	rootCmd.AddCommand(logoutCmd)

	// Users
	usersCmd := &cobra.Command{
		Use:   "users",
		Short: "Gérer les utilisateurs",
	}
	rootCmd.AddCommand(usersCmd)

	// Users List
	listUsersCmd := &cobra.Command{
		Use:   "list",
		Short: "Lister tous les utilisateurs",
		Run:   listUsers,
	}
	usersCmd.AddCommand(listUsersCmd)

	// Users Get
	getUserCmd := &cobra.Command{
		Use:   "get [user_id]",
		Short: "Récupérer un utilisateur spécifique",
		Args:  cobra.ExactArgs(1),
		Run:   getUser,
	}
	usersCmd.AddCommand(getUserCmd)

	// Users Create
	createUserCmd := &cobra.Command{
		Use:   "create",
		Short: "Créer un nouvel utilisateur",
		Run:   createUser,
	}
	createUserCmd.Flags().String("name", "", "Le nom de l'utilisateur")
	createUserCmd.Flags().String("email", "", "L'adresse email de l'utilisateur")
	createUserCmd.Flags().String("password", "", "Le mot de passe de l'utilisateur")
	createUserCmd.Flags().StringSlice("roles", nil, "Les rôles de l'utilisateur (potentiellement vide)")
	createUserCmd.Flags().StringSlice("groups", nil, "Les groupes de l'utilisateur (potentiellement vide)")
	usersCmd.AddCommand(createUserCmd)

	// Users Update
	updateUserCmd := &cobra.Command{
		Use:   "update [user_id]",
		Short: "Mettre à jour un utilisateur existant",
		Args:  cobra.ExactArgs(1),
		Run:   updateUser,
	}
	updateUserCmd.Flags().String("name", "", "Le nouveau nom de l'utilisateur")
	updateUserCmd.Flags().String("email", "", "La nouvelle adresse email de l'utilisateur")
	updateUserCmd.Flags().String("password", "", "Le nouveau mot de passe de l'utilisateur")
	updateUserCmd.Flags().StringSlice("roles", nil, "Les nouveaux rôles de l'utilisateur")
	updateUserCmd.Flags().StringSlice("groups", nil, "Les nouveaux groupes de l'utilisateur")
	usersCmd.AddCommand(updateUserCmd)

	// Users Delete
	deleteUserCmd := &cobra.Command{
		Use:   "delete [user_id]",
		Short: "Supprimer un utilisateur existant",
		Args:  cobra.ExactArgs(1),
		Run:   deleteUser,
	}
	usersCmd.AddCommand(deleteUserCmd)

	// Roles
	rolesCmd := &cobra.Command{
		Use:   "roles",
		Short: "Gérer les rôles",
	}
	rootCmd.AddCommand(rolesCmd)

	// Roles List
	listRolesCmd := &cobra.Command{
		Use:   "list",
		Short: "Lister tous les rôles",
		Run:   listRoles,
	}
	rolesCmd.AddCommand(listRolesCmd)

	// Roles Get
	getRoleCmd := &cobra.Command{
		Use:   "get [role_id]",
		Short: "Récupérer un rôle spécifique",
		Args:  cobra.ExactArgs(1),
		Run:   getRole,
	}
	rolesCmd.AddCommand(getRoleCmd)

	// Roles Create
	createRoleCmd := &cobra.Command{
		Use:   "create",
		Short: "Créer un nouveau rôle",
		Run:   createRole,
	}
	createRoleCmd.Flags().String("name", "", "Le nom du rôle")
	createRoleCmd.Flags().String("description", "", "La description du rôle")
	rolesCmd.AddCommand(createRoleCmd)

	// Roles Update
	updateRoleCmd := &cobra.Command{
		Use:   "update [role_id]",
		Short: "Mettre à jour un rôle existant",
		Args:  cobra.ExactArgs(1),
		Run:   updateRole,
	}
	updateRoleCmd.Flags().String("name", "", "Le nouveau nom du rôle")
	updateRoleCmd.Flags().String("description", "", "La nouvelle description du rôle")
	rolesCmd.AddCommand(updateRoleCmd)

	// Roles Delete
	deleteRoleCmd := &cobra.Command{
		Use:   "delete [role_id]",
		Short: "Supprimer un rôle existant",
		Args:  cobra.ExactArgs(1),
		Run:   deleteRole,
	}
	rolesCmd.AddCommand(deleteRoleCmd)

	// Groups
	groupsCmd := &cobra.Command{
		Use:   "groups",
		Short: "Gérer les groupes d'utilisateurs",
	}
	rootCmd.AddCommand(groupsCmd)

	// Groups List
	listGroupsCmd := &cobra.Command{
		Use:   "list",
		Short: "Lister tous les groupes",
		Run:   listGroups,
	}
	groupsCmd.AddCommand(listGroupsCmd)

	// Groups Get
	getGroupCmd := &cobra.Command{
		Use:   "get [group_id]",
		Short: "Récupérer un groupe spécifique",
		Args:  cobra.ExactArgs(1),
		Run:   getGroup,
	}
	groupsCmd.AddCommand(getGroupCmd)

	// Groups Create
	createGroupCmd := &cobra.Command{
		Use:   "create",
		Short: "Créer un nouveau groupe",
		Run:   createGroup,
	}
	createGroupCmd.Flags().String("name", "", "Le nom du groupe")
	createGroupCmd.Flags().String("parent_group_id", "", "L'ID du groupe parent")
	groupsCmd.AddCommand(createGroupCmd)

	// Groups Update
	updateGroupCmd := &cobra.Command{
		Use:   "update [group_id]",
		Short: "Mettre à jour un groupe existant",
		Args:  cobra.ExactArgs(1),
		Run:   updateGroup,
	}

	// Roles Delete
	deleteGroupCmd := &cobra.Command{
		Use:   "delete [group_id]",
		Short: "Supprimer un groupe existant",
		Args:  cobra.ExactArgs(1),
		Run:   deleteGroup,
	}
	rolesCmd.AddCommand(deleteGroupCmd)

	updateGroupCmd.Flags().String("name", "", "Le nouveau nom du groupe")
	groupsCmd.AddCommand(updateGroupCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

////////////////////////////////////////////////////////////////	//////////////////////////////////////////////

func sendRequest(method, url string, headers map[string]string, body []byte) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func login(cmd *cobra.Command, args []string) {
	email, _ := cmd.Flags().GetString("email")
	password, _ := cmd.Flags().GetString("password")

	payload := map[string]string{
		"email":    email,
		"password": password,
	}
	jsonPayload, _ := json.Marshal(payload)

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	responseBody, err := sendRequest("POST", "http://app:8080/login", headers, jsonPayload)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(string(responseBody))
}

func serve(cmd *cobra.Command, args []string) {
	for {
		time.Sleep(time.Hour)
	}
}

func refresh(cmd *cobra.Command, args []string) {
	refreshToken, _ := cmd.Flags().GetString("refresh_token")

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", refreshToken),
	}
	responseBody, err := sendRequest("POST", "http://app:8080/refresh", headers, nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(string(responseBody))
}

func logout(cmd *cobra.Command, args []string) {
	accessToken, _ := cmd.Flags().GetString("access_token")
	refreshToken, _ := cmd.Flags().GetString("refresh_token")

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", accessToken),
	}
	responseBody, err := sendRequest("DELETE", fmt.Sprintf("http://app:8080/logout/%s", refreshToken), headers, nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(string(responseBody))
	fmt.Println("Successfully logged out")
}

func listUsers(cmd *cobra.Command, args []string) {
	responseBody, err := sendRequest("GET", "http://app:8080/users/", nil, nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(string(responseBody))
}

func updateUser(cmd *cobra.Command, args []string) {
	userId := args[0]
	email, _ := cmd.Flags().GetString("email")
	password, _ := cmd.Flags().GetString("password")
	name, _ := cmd.Flags().GetString("name")

	payload := map[string]string{
		"email":    email,
		"password": password,
		"name":     name,
	}
	jsonPayload, _ := json.Marshal(payload)

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	responseBody, err := sendRequest("PUT", fmt.Sprintf("http://app:8080/users/%s", userId), headers, jsonPayload)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(string(responseBody))
}

func getUser(cmd *cobra.Command, args []string) {
	userId := args[0]
	responseBody, err := sendRequest("GET", fmt.Sprintf("http://app:8080/users/%s", userId), nil, nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(string(responseBody))
}

func createUser(cmd *cobra.Command, args []string) {
	email, _ := cmd.Flags().GetString("email")
	password, _ := cmd.Flags().GetString("password")
	name, _ := cmd.Flags().GetString("name")

	payload := map[string]string{
		"email":    email,
		"password": password,
		"name":     name,
	}
	jsonPayload, _ := json.Marshal(payload)

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	responseBody, err := sendRequest("POST", "http://app:8080/users/", headers, jsonPayload)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(string(responseBody))
}

func deleteUser(cmd *cobra.Command, args []string) {
	userId := args[0]
	responseBody, err := sendRequest("DELETE", fmt.Sprintf("http://app:8080/users/%s", userId), nil, nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(string(responseBody))
}

func listRoles(cmd *cobra.Command, args []string) {
	responseBody, err := sendRequest("GET", "http://app:8080/roles/", nil, nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(string(responseBody))
}

func getRole(cmd *cobra.Command, args []string) {
	roleId := args[0]
	responseBody, err := sendRequest("GET", fmt.Sprintf("http://app:8080/roles/%s", roleId), nil, nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(string(responseBody))
}

func createRole(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	description, _ := cmd.Flags().GetString("description")

	payload := map[string]string{
		"name":        name,
		"description": description,
	}
	jsonPayload, _ := json.Marshal(payload)

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	responseBody, err := sendRequest("POST", "http://app:8080/roles", headers, jsonPayload)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(string(responseBody))
}

func updateRole(cmd *cobra.Command, args []string) {
	roleID, _ := cmd.Flags().GetString("id")
	name, _ := cmd.Flags().GetString("name")
	description, _ := cmd.Flags().GetString("description")

	payload := map[string]string{
		"name":        name,
		"description": description,
	}
	jsonPayload, _ := json.Marshal(payload)

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	endpoint := fmt.Sprintf("http://app:8080/roles/%s", roleID)
	responseBody, err := sendRequest("PUT", endpoint, headers, jsonPayload)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(string(responseBody))
}

func deleteRole(cmd *cobra.Command, args []string) {
	roleId := args[0]
	responseBody, err := sendRequest("DELETE", fmt.Sprintf("http://app:8080/roles/%s", roleId), nil, nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(string(responseBody))
}

func listGroups(cmd *cobra.Command, args []string) {
	responseBody, err := sendRequest("GET", "http://app:8080/groups/", nil, nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(string(responseBody))
}

func getGroup(cmd *cobra.Command, args []string) {
	groupId := args[0]
	responseBody, err := sendRequest("GET", fmt.Sprintf("http://app:8080/roles/%s", groupId), nil, nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(string(responseBody))
}

func deleteGroup(cmd *cobra.Command, args []string) {
	groupId := args[0]
	responseBody, err := sendRequest("DELETE", fmt.Sprintf("http://app:8080/roles/%s", groupId), nil, nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(string(responseBody))
}

func createGroup(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	description, _ := cmd.Flags().GetString("description")

	payload := map[string]string{
		"name":        name,
		"description": description,
	}
	jsonPayload, _ := json.Marshal(payload)

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	responseBody, err := sendRequest("POST", "http://app:8080/groups", headers, jsonPayload)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(string(responseBody))
}

func updateGroup(cmd *cobra.Command, args []string) {
	groupID, _ := cmd.Flags().GetString("id")
	name, _ := cmd.Flags().GetString("name")
	description, _ := cmd.Flags().GetString("description")

	payload := map[string]string{
		"name":        name,
		"description": description,
	}
	jsonPayload, _ := json.Marshal(payload)

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	endpoint := fmt.Sprintf("http://app:8080/groups/%s", groupID)
	responseBody, err := sendRequest("PUT", endpoint, headers, jsonPayload)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(string(responseBody))
}
