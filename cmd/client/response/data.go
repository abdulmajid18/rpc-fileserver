package response

type FileResponse struct {
	GenericResponse
	Contents []byte
}

type FileMetadataResponse struct {
	Success     bool
	Message     string
	Size        int64
	Mode        string
	ModifiedAt  string
	IsDirectory bool
}

type GenericResponse struct {
	Success bool
	Message string
}
