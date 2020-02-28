package main

func (s *server) routes() {
	s.router.HandleFunc("/admin/", s.handleAdmin).Methods("GET")
	s.router.HandleFunc("/api", s.checkAPI).Methods("GET")
}
