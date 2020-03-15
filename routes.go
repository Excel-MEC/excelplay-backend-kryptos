package main

func (s *server) routes() {
	s.router.HandleFunc("/admin/", s.handleAdmin).Methods("GET")
	s.router.Handle("/api/ping", s.heartBeat()).Methods("GET")
	s.router.Handle("/api/question", s.authMiddleware(s.handleNextQuestion())).Methods("GET")
	s.router.Handle("/api/submit", s.authMiddleware(s.handleSubmission())).Methods("POST")
}
