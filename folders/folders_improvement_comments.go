package folders

/*
import (
	"github.com/gofrs/uuid"
)

// GetAllFoldersAnnotated This function performs the same task as FetchAllFoldersByOrgID, which means the function name
// is misleading. It can be renamed, and redundant logic in FetchAllFoldersByOrgID can be removed. There are also non-descriptive
// variable names that can be improved.
func GetAllFoldersAnnotated(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// Since these are currently redundant, use them or remove them
	var (
		err error
		f1  Folder
		fs  []*Folder
	)
	f := []Folder{}
	// The error is discarded when it should be returned
	r, _ := FetchAllFoldersByOrgIDAnnotated(req.OrgID)
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
*/
