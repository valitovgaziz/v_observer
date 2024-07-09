package apiserver

// APIServer ...
type APIServer struct {}


// New ...
func New() *APIServer {
	return new(APIServer)
}


// Start ...
func (s *APIServer) Start() error {
	return nil
}