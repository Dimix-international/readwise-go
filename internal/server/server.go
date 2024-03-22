package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Dimix-international/readwise-go/db"
	"github.com/Dimix-international/readwise-go/internal/config"
	"github.com/Dimix-international/readwise-go/internal/handlers"
	"github.com/Dimix-international/readwise-go/internal/models"
	"github.com/Dimix-international/readwise-go/internal/service"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Server struct {
	cfg     config.Config
	closers []models.CloseFunc
}

func NewServer(cfg config.Config) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Run() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := s.launchSever(); err != nil {
			log.Printf("error closing server: %v", err)
			exit <- syscall.SIGTERM
			close(exit)
		}
	}()

	<-exit

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := s.Shutdown(shutdownCtx); err != nil {
		log.Printf("error closing server: %v", err)
	}
}

func (s *Server) launchSever() error {
	client, err := db.NewDb(mysql.Config{User: "root", Passwd: "dima94dimix@MySQL",  Net: "tcp", DBName: "readwise", Addr: "localhost:3306"})
	if err != nil {
		return err
	}

	router := mux.NewRouter()

	httpServer := &http.Server{
		Handler:      router,
		Addr:         s.cfg.HTTPServer.Address,
		ReadTimeout:  s.cfg.HTTPServer.Timeout,
		WriteTimeout: s.cfg.HTTPServer.Timeout,
		IdleTimeout:  s.cfg.HTTPServer.IdleTimeout,
	}

	handlers.NewFileHandler(router, service.NewFileService(db.NewBookStorage(client.DB))).RegisterRoutes()

	log.Printf("server started on port: %v", s.cfg.HTTPServer.Port)

	if err := httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (s *Server) AddCloser(closer models.CloseFunc) {
	s.closers = append(s.closers, closer)
}

func (s *Server) Shutdown(ctx context.Context) error {
	for _, fn := range s.closers {
		if err := fn(ctx); err != nil {
			return err
		}
	}

	return nil
}
