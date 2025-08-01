package fileservice

type WriteFileArgs struct {
	FileID  uintptr   // File descriptor or identifier
	Content []byte    // Data to write
	Mode    WriteMode // SingleWrite or ChunkedWrite
	Offset  int64     // Required for chunked writes (last position)
}

type WriteMode string

const (
	ModeSingleWrite  WriteMode = "single"  // Write entire content at once
	ModeChunkedWrite WriteMode = "chunked" // Append content at Offset
)

type WriteFileReply struct {
	Success      bool
	BytesWritten int
	NextOffset   int64 // For chunked: returns new file end position
}
