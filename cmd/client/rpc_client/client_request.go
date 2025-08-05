package rpcclient

import "abdulmajid/fileserver/internal/types"

func NewDirRequest(name string) *types.DirRequest {
	return &types.DirRequest{Name: name}
}
