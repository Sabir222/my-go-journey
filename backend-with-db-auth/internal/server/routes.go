package server

func (s *FiberServer) RegisterRoutes() {
	s.App.Get("/", s.HelloWorldHandler)
	s.App.Get("/health", s.healthHandler)
	s.App.Post("/register", s.RegisterHandler)
}
