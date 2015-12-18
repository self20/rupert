package rupert

import (
	"crypto/tls"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewEngine() *gin.Engine {
	engine := gin.New()

	// Setup middlewares
	engine.Use(gin.Logger())

	// User handlers
	engine.POST("/user", user_create)
	engine.DELETE("/user/:user_id", user_delete)
	engine.GET("/user/:user_id", user_get)
	engine.POST("/user/authenticate", user_authenticate)
	engine.GET("/users", users_get)

	// Torrent handlers
	engine.POST("/torrent", torrent_create)
	engine.DELETE("/torrent/:torrent_id", torrent_delete)
	engine.GET("/torrent/:torrent_id", torrent_get)
	engine.PATCH("/torrent/:torrent_id", torrent_update)

	// Wiki handlers
	engine.POST("/wiki/:name", wiki_update)
	engine.GET("/wiki/:name", wiki_get)
	engine.DELETE("/wiki/:name", wiki_delete)
	engine.GET("/wikis", wiki_list)

	// Forum handlers
	engine.POST("/forums", forum_create)
	engine.GET("/forums", forums_list_get)
	engine.DELETE("/forums/:forum_id", forum_delete)
	engine.PATCH("/forums/:forum_id", forum_update)
	engine.POST("/forums/:forum_id", forum_thread_create)
	engine.GET("/forums/:forum_id", forum_get)
	engine.GET("/forums/:forum_id/threads", forum_threads_get)
	engine.DELETE("/threads/:thread_id", forum_thread_delete)
	engine.PATCH("/threads/:forum_id", forum_thread_update)
	engine.POST("/threads/:thread_id", forum_comment_create)
	engine.DELETE("/comments/:comment_id", forum_comment_delete)
	engine.PATCH("/comments/:comment_id", forum_comment_update)
	engine.GET("/threads/:thread_id/comments", forum_thread_comments_get)

	return engine
}

func ListenAndServe(engine *gin.Engine) {
	if config.SSLCert == "" || config.SSLPrivateKey == "" {
		// Non-TLS enabled API
		log.Fatalln("SSL config keys not set in config!")
		if !config.Debug {
			log.Warnln("Running in production without TLS enforced is highly unrecommended")
		}
		srv := http.Server{Addr: config.ListenHost, Handler: engine}
		srv.ListenAndServe()
	} else {
		tls_config := &tls.Config{
			MinVersion:               tls.VersionTLS12,
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
			},
		}
		srv := http.Server{TLSConfig: tls_config, Addr: config.ListenHost, Handler: engine}
		srv.ListenAndServeTLS(config.SSLCert, config.SSLPrivateKey)
	}
}
