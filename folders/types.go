package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID    uuid.UUID
	PageSize int
	Token    string
}

type FetchFolderResponse struct {
	Folders []*Folder
	Token   string
}
