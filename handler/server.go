package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/monirz/gojwt"
	"github.com/monirz/gojwt/config"
	"github.com/monirz/gojwt/middleware"

	"github.com/go-chi/chi"
)

type Server struct {
	router      *chi.Mux
	Config      *config.Config
	UserService gojwt.UserService
}

func NewServer() *Server {

	// var err error
	s := &Server{}
	s.Config = config.NewConfig()

	s.router = chi.NewRouter()

	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome")
	})

	s.router.Route("/api/v1", func(r chi.Router) {
		r.Post("/login", s.Login)
		r.Post("/signup", s.CreateUserHandler)

		r.Group(func(r chi.Router) {
			r.Use(middleware.JWTAuth)
		})

	})

	return s

}

func (s *Server) Run() {

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	addr := fmt.Sprintf(":%s", s.Config.Port)

	srv := &http.Server{
		Addr: addr,
		// ReadTimeout:  60 * time.Second,
		// WriteTimeout: 60 * time.Second,
		Handler: s.router,
	}

	go func() {
		log.Println("Staring server with address ", addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Failed to start http server on :", err)
			os.Exit(-1)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Http server couldn't shutdown gracefully")
	}

	log.Println("shutting down")

}
