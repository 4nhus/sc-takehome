package folders

import (
	"github.com/gofrs/uuid"
)

// GetAllFoldersAnnotated This function performs the same task as FetchAllFoldersByOrgID, so one of them can be removed
func GetAllFoldersAnnotated(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// Since these are currently redundant, use them or remove them
	var (
		err error
		f1  Folder
		f1  Folder
		fs  []*Folder
	)
	f := []Folder{}
	// The error is discarded when it should be returned
	r, _ := FetchAllFoldersByOrgID(req.OrgID)
	// Dereferencing each folder pointer is redundant as the FetchFolderResponse value wraps a slice of folder pointers
	for k, v := range r {
		f = append(f, *v)
	}
	var fp []*Folder
	for k1, v1 := range f {
		fp = append(fp, &v1)
	}
	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp}
	return ffr, nil
}

func FetchAllFoldersByOrgIDAnnotated(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
