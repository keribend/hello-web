package handlers

type FileHandler struct {
	dir string
}

func NewFileHandler(path string) FileHandler {
	return FileHandler{path}
}
