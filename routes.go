package main

func (s *server) routes() {
	s.router.HandleFunc("/admin/", s.handleAdmin).Methods("GET")
	s.router.Handle("/api/ping", httpHandler(s.heartBeat())).Methods("GET")
	s.router.Handle("/api/question", s.authMiddleware(httpHandler(s.handleNextQuestion()))).Methods("GET")
	s.router.Handle("/api/submit", s.authMiddleware(httpHandler(s.handleSubmission()))).Methods("POST")
}
