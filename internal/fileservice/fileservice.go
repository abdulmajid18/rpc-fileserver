package fileservice

import (
	"abdulmajid/fileserver/internal/types"
	"os"
	"time"
)

type FileService interface {
	CreateDir(req types.DirRequest, res *types.GenericResponse) error
	CreateFile(req types.FileRequest, res *types.GenericResponse) error
	ReadFile(req types.FileRequest, res *types.FileResponse) error
	WriteFile(req types.FileRequest, res *types.GenericResponse) error
	AppendFile(req types.FileRequest, res *types.GenericResponse) error
	GetFileInfo(req types.FileRequest, res *types.FileMetadataResponse) error
}

type FileOperations struct{}

func (f *FileOperations) CreateDir(req types.DirRequest, res *types.GenericResponse) error {
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

func (f *FileOperations) CreateFile(req types.FileRequest, res *types.GenericResponse) error {
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

func (f *FileOperations) ReadFile(req types.FileRequest, res *types.FileResponse) error {
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

func (f *FileOperations) WriteFile(req types.FileRequest, res *types.GenericResponse) error {
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

func (f *FileOperations) AppendFile(req types.FileRequest, res *types.GenericResponse) error {
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

func (f *FileOperations) GetFileInfo(req types.FileRequest, res *types.FileMetadataResponse) error {
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
