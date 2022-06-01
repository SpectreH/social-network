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
	mux.Handle("/api/profile/removefollow", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.ProfileRemoveFollow)))
	mux.Handle("/api/profile/unfollow", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.ProfileUnFollow)))
	mux.Handle("/api/profile/requesttofollow", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.ProfileRequestToFollow)))

	mux.Handle("/api/socket", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.CreateSocketReader)))
	mux.Handle("/api/requests", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.GetRequestList)))
	mux.Handle("/api/acceptrequest", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.AcceptRequest)))
	mux.Handle("/api/declinerequest", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.DeclineRequest)))

	mux.Handle("/api/post/new", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.PostNew)))
	mux.Handle("/api/post", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.GetPostContent)))

	mux.Handle("/api/followers", middleware.Repo.SetupCors(http.HandlerFunc(handlers.Repo.FollowerList)))

	fileServer := http.FileServer(http.Dir("./images"))
	mux.Handle("/images/", http.StripPrefix("/images", fileServer))

	return mux
}
