package request

type FileRequest struct {
	Filename string
	Contents []byte
}

type DirRequest struct {
	Name string
}

type FileResponse struct {
	Contents []byte
}
