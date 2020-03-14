package main

func (s *server) routes() {
	s.router.HandleFunc("/admin/", s.handleAdmin).Methods("GET")
	s.router.HandleFunc("/api/ping", s.heartBeat).Methods("GET")
	s.router.Handle("/api/question", httpHandler(s.handleNextQuestion())).Methods("GET")
	s.router.Handle("/api/submit", httpHandler(s.handleSubmission())).Methods("POST")
}
