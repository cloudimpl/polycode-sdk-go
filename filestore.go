package polycode

type File interface {
	Parent() Folder
	Name() string
	Path() string

	Get() (bool, []byte, error)
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

type Folder interface {
	Parent() Folder
	Name() string
	Path() string

	Folder(name string) Folder
	CreateNewFolder(name string) (Folder, error)
	File(name string) File
}
