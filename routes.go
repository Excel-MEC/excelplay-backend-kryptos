package main

func (s *server) routes() {
	s.router.HandleFunc("/admin/", s.handleAdmin).Methods("GET")
	s.router.HandleFunc("/api/ping", s.heartBeat).Methods("GET")
	s.router.HandleFunc("/api/question", s.handleNextQuestion()).Methods("GET")
	s.router.HandleFunc("/api/submit", s.handleSubmission()).Methods("POST")

	s.router.Use(s.authMiddleware)
}
