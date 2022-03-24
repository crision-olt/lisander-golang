package controller

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/crision98/lisander-golang-backend/middleware"
	"github.com/crision98/lisander-golang-backend/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Controllers set the port, the handler and puts the server to listen*/
func Controllers() {
	router := mux.NewRouter()
	router.HandleFunc("/notification", middleware.CheckDB(middleware.ValidateJWT(routes.GetNotifications))).Methods("GET")
	router.HandleFunc("/notification", middleware.CheckDB(middleware.ValidateJWT(routes.CheckNotification))).Methods("DELETE")
	router.HandleFunc("/count/notification", middleware.CheckDB(middleware.ValidateJWT(routes.CountNotifications))).Methods("GET")

	router.HandleFunc("/register", middleware.CheckDB(routes.InsertUser)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(routes.Login)).Methods("POST")

	router.HandleFunc("/admin", middleware.CheckDB(middleware.ValidateJWT(routes.GetAdmins))).Methods("GET")
	router.HandleFunc("/admin", middleware.CheckDB(middleware.ValidateJWT(routes.InsertAdmin))).Methods("POST")
	router.HandleFunc("/admin", middleware.CheckDB(middleware.ValidateJWT(routes.BanAdmin))).Methods("DELETE")

	router.HandleFunc("/user/changePassword", middleware.CheckDB(middleware.ValidateJWT(routes.ChangePassword))).Methods("POST")
	router.HandleFunc("/user", middleware.CheckDB(middleware.ValidateJWT(routes.GetUser))).Methods("GET")
	router.HandleFunc("/user", middleware.CheckDB(middleware.ValidateJWT(routes.ModifyUser))).Methods("POST")
	router.HandleFunc("/user/avatar", middleware.CheckDB(middleware.ValidateJWT(routes.UploadAvatar))).Methods("POST")
	router.HandleFunc("/user/banner", middleware.CheckDB(middleware.ValidateJWT(routes.UploadBanner))).Methods("POST")

	router.HandleFunc("/toot", middleware.CheckDB(middleware.ValidateJWT(routes.GetToot))).Methods("GET")
	router.HandleFunc("/toots", middleware.CheckDB(middleware.ValidateJWT(routes.InsertToot))).Methods("POST")
	router.HandleFunc("/toots", middleware.CheckDB(middleware.ValidateJWT(routes.GetTootsFromUser))).Methods("GET")
	router.HandleFunc("/toots", middleware.CheckDB(middleware.ValidateJWT(routes.DeleteToot))).Methods("DELETE")
	router.HandleFunc("/count/toots", middleware.CheckDB(middleware.ValidateJWT(routes.CountTootsFromUser))).Methods("GET")

	router.HandleFunc("/comments", middleware.CheckDB(middleware.ValidateJWT(routes.InsertComment))).Methods("POST")
	router.HandleFunc("/comments", middleware.CheckDB(middleware.ValidateJWT(routes.GetCommentsFromToot))).Methods("GET")
	router.HandleFunc("/comments", middleware.CheckDB(middleware.ValidateJWT(routes.DeleteComment))).Methods("DELETE")
	router.HandleFunc("/count/comments", middleware.CheckDB(middleware.ValidateJWT(routes.CountCommentsFromToot))).Methods("GET")

	router.HandleFunc("/report", middleware.CheckDB(middleware.ValidateJWT(routes.GetReport))).Methods("GET")
	router.HandleFunc("/reports", middleware.CheckDB(middleware.ValidateJWT(routes.InsertUserReport))).Methods("POST")
	router.HandleFunc("/reports", middleware.CheckDB(middleware.ValidateJWT(routes.ValidateUserReport))).Methods("DELETE")
	router.HandleFunc("/reports", middleware.CheckDB(middleware.ValidateJWT(routes.GetReports))).Methods("GET")

	router.HandleFunc("/relation", middleware.CheckDB(middleware.ValidateJWT(routes.InsertRelation))).Methods("POST")
	router.HandleFunc("/relation", middleware.CheckDB(middleware.ValidateJWT(routes.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/relation", middleware.CheckDB(middleware.ValidateJWT(routes.CheckRelation))).Methods("GET")
	router.HandleFunc("/count/relations", middleware.CheckDB(middleware.ValidateJWT(routes.CountRelations))).Methods("GET")

	router.HandleFunc("/users/relations", middleware.CheckDB(middleware.ValidateJWT(routes.GetUsersFromRelations))).Methods("GET")
	router.HandleFunc("/users/blocks", middleware.CheckDB(middleware.ValidateJWT(routes.GetBlockedUsers))).Methods("GET")
	router.HandleFunc("/users", middleware.CheckDB(middleware.ValidateJWT(routes.FindUsersByEmail))).Methods("GET")

	router.HandleFunc("/users/toots", middleware.CheckDB(middleware.ValidateJWT(routes.GetTootsFromRelations))).Methods("GET")

	router.HandleFunc("/isStandard", middleware.CheckDB(middleware.ValidateJWT(routes.IsStandard))).Methods("GET")
	router.HandleFunc("/isSuperAdmin", middleware.CheckDB(middleware.ValidateJWT(routes.IsSuperAdmin))).Methods("GET")
	router.HandleFunc("/isAdmin", middleware.CheckDB(middleware.ValidateJWT(routes.IsAdmin))).Methods("GET")
	router.HandleFunc("/reloadDate", middleware.CheckDB(middleware.ValidateJWT(routes.ReloadDate))).Methods("GET")
	router.HandleFunc("/canReloadDate", middleware.CheckDB(middleware.ValidateJWT(routes.CanReloadDate))).Methods("GET")
	router.HandleFunc("/date", middleware.CheckDB(middleware.ValidateJWT(routes.GetDate))).Methods("GET")
	router.HandleFunc("/reloads", middleware.CheckDB(middleware.ValidateJWT(routes.GetReloads))).Methods("GET")

	router.HandleFunc("/block", middleware.CheckDB(middleware.ValidateJWT(routes.InsertBlock))).Methods("POST")
	router.HandleFunc("/block", middleware.CheckDB(middleware.ValidateJWT(routes.DeleteBlock))).Methods("DELETE")
	router.HandleFunc("/block", middleware.CheckDB(middleware.ValidateJWT(routes.CheckBlock))).Methods("GET")

	router.PathPrefix("/uploads/avatars/").Handler(http.StripPrefix("/uploads/avatars/", http.FileServer(http.Dir("."+"/uploads/avatars/"))))
	router.PathPrefix("/uploads/banners/").Handler(http.StripPrefix("/uploads/banners/", http.FileServer(http.Dir("."+"/uploads/banners/"))))

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	controller := cors.AllowAll().Handler(router)
	srv := &http.Server{
		Addr:         ":3000",
		Handler:      controller,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Fatal(srv.ListenAndServeTLS("tls.crt", "tls.key"))
	//log.Fatal(srv.ListenAndServe())
}
