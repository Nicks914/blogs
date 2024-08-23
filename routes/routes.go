package routes

import (
	"blogs/controllers"
	"net/http"
)

func InitializeRoutes() {
	http.HandleFunc("/api/register", controllers.Register)
	http.HandleFunc("/api/login", controllers.Login)

	http.HandleFunc("/api/profile", controllers.GetProfile)
	http.HandleFunc("/api/profile/update", controllers.UpdateProfile)

	http.HandleFunc("/api/posts", controllers.GetPosts)
	http.HandleFunc("/api/posts/create", controllers.CreatePost)
	http.HandleFunc("/api/posts/update", controllers.UpdatePost)
	http.HandleFunc("/api/posts/delete", controllers.DeletePost)
}
