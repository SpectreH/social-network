package routes

import (
	"net/http"
	"social-network/internal/handlers"

	middleware "social-network/internal/middleware"
)

// SetRoutes sets handler and load server files
func SetRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/api/logout", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.LogOut)))
	mux.Handle("/api/signup", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.SignUp)))
	mux.Handle("/api/signin", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.SignIn)))
	mux.Handle("/api/authme", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.AuthMe)))

	mux.Handle("/api/settings/updateAvatar", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.UpdateAvatar)))
	mux.Handle("/api/settings/updateProfile", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.UpdateProfile)))
	mux.Handle("/api/settings/updatePrivacy", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.UpdatePrivacy)))

	mux.Handle("/api/profile/fetchProfile", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.GetProfile)))
	mux.Handle("/api/profile/follow", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.ProfileFollow)))
	mux.Handle("/api/profile/unfollow", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.ProfileUnFollow)))
	mux.Handle("/api/profile/requesttofollow", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.ProfileRequestToFollow)))

	fileServer := http.FileServer(http.Dir("./images"))
	mux.Handle("/images/", http.StripPrefix("/images", fileServer))

	return mux
}
