package fileservice

import (
	"abdulmajid/fileserver/cmd/server/request"
	"abdulmajid/fileserver/cmd/server/response"
	"os"
	"time"
)

type FileService interface {
	CreateDir(req request.DirRequest, res *response.GenericResponse) error
	CreateFile(req request.FileRequest, res *response.GenericResponse) error
	ReadFile(req request.FileRequest, res *response.FileResponse) error
	WriteFile(req request.FileRequest, res *response.GenericResponse) error
	AppendFile(req request.FileRequest, res *response.GenericResponse) error
	GetFileInfo(req request.FileRequest, res *response.FileMetadataResponse) error
}

type FileOperations struct{}

func (f *FileOperations) CreateDir(req request.DirRequest, res *response.GenericResponse) error {
	err := os.MkdirAll(req.Name, os.ModePerm)
	if err != nil {
		res.Success = false
		res.Message = err.Error()
		return nil
	}
	res.Success = true
	res.Message = "Directory created"
	return nil
}

func (f *FileOperations) CreateFile(req request.FileRequest, res *response.GenericResponse) error {
	file, err := os.Create(req.Filename)
	if err != nil {
		res.Success = false
		res.Message = "Failed to create file: " + err.Error()
		return nil
	}
	defer file.Close()

	res.Success = true
	res.Message = "File created successfully"
	return nil
}

func (f *FileOperations) ReadFile(req request.FileRequest, res *response.FileResponse) error {
	data, err := os.ReadFile(req.Filename)
	if err != nil {
		res.Success = false
		res.Message = err.Error()
		return nil
	}
	res.Success = true
	res.Message = "File read"
	res.Contents = data
	return nil
}

func (f *FileOperations) WriteFile(req request.FileRequest, res *response.GenericResponse) error {
	err := os.WriteFile(req.Filename, req.Contents, 0644)
	if err != nil {
		res.Success = false
		res.Message = err.Error()
		return nil
	}
	res.Success = true
	res.Message = "File written"
	return nil
}

func (f *FileOperations) AppendFile(req request.FileRequest, res *response.GenericResponse) error {
	file, err := os.OpenFile(req.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		res.Success = false
		res.Message = "Failed to open file: " + err.Error()
		return nil
	}
	defer file.Close()

	_, err = file.Write(req.Contents)
	if err != nil {
		res.Success = false
		res.Message = "Failed to write to file: " + err.Error()
		return nil
	}

	res.Success = true
	res.Message = "Content appended successfully"
	return nil
}

func (f *FileOperations) GetFileInfo(req request.FileRequest, res *response.FileMetadataResponse) error {
	info, err := os.Stat(req.Filename)
	if err != nil {
		res.Success = false
		res.Message = "Failed to get file info: " + err.Error()
		return nil
	}

	res.Success = true
	res.Message = "File info retrieved"
	res.Size = info.Size()
	res.Mode = info.Mode().String()
	res.ModifiedAt = info.ModTime().Format(time.RFC3339)
	res.IsDirectory = info.IsDir()
	return nil
}
