package main

import (
	"fmt"
	"net/http"
	"shop/config"
	"shop/router"
)

func main() {
	handler := router.InitRouter()
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", config.Port),
		Handler: handler,
		ReadTimeout: config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		return 
	}
}
