package polycode

import "time"

type FileMetaData struct {
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Size     int64     `json:"size"`
}

type ReadOnlyFile interface {
	Parent() ReadOnlyFolder
	Name() string
	Path() string
	Metadata() FileMetaData

	Read() ([]byte, error)
	Download(localFilePath string) error
	GetDownloadLink() (string, error)
}

type File interface {
	Parent() Folder
	Name() string
	Path() string
	Metadata() FileMetaData

	Read() ([]byte, error)
	Download(filePath string) error
	GetDownloadLink() (string, error)

	Save(data []byte) error
	Upload(filePath string) error
	GetUploadLink() (string, error)

	Delete() error
	Rename(newName string) error
	MoveTo(dest Folder) error
	CopyTo(dest Folder) error
}

type ReadOnlyFolder interface {
	Parent() ReadOnlyFolder
	Name() string
	Path() string

	Folder(name string) ReadOnlyFolder
	File(name string) ReadOnlyFile
	List(maxFiles int32, offsetToken *string) ([]ReadOnlyFile, *string, error)
}

type Folder interface {
	Parent() Folder
	Name() string
	Path() string

	Folder(name string) Folder
	CreateNewFolder(name string) (Folder, error)
	File(name string) File
	List(maxFiles int32, offsetToken *string) ([]File, *string, error)
}
