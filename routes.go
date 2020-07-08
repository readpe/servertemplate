package main

// server routes, assigned to server router
func (s *server) routes() {
	s.router.HandleFunc("/", s.handleIndex())
	s.router.HandleFunc("/about", s.handleAbout())
	s.router.HandleFunc("/admin/", s.adminOnly(s.handleIndex()))
}
