package types

type FileRequest struct {
	Filename string
	Contents []byte
}

type DirRequest struct {
	Name string
}
