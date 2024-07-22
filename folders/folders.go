package folders

import (
	"github.com/gofrs/uuid"
)

// GetAllFolders retrieves all folders for a given orgID.
// It takes a FetchFolderRequest - which wraps an orgID - and returns a FetchFolderResponse - which wraps a slice of
// folder pointers - and an error, if any.
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// Declares (unused) variables
	var (
		err error
		f1  Folder
		fs  []*Folder
	)
	// Initialises f as an empty slice of folders
	f := []Folder{}
	// Fetches all folders with the same orgID as in req, ignoring the error
	r, _ := FetchAllFoldersByOrgID(req.OrgID)
	// Appends each fetched folder to f
	for k, v := range r {
		f = append(f, *v)
	}
	// Initialises fp as an empty slice of folder pointers
	var fp []*Folder
	// Appends the pointer to each folder in f to fp
	for k1, v1 := range f {
		fp = append(fp, &v1)
	}
	// Initialises ffr as the FetchFolderResponse wrapping fp
	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp}
	return ffr, nil
}

// FetchAllFoldersByOrgID fetches all folders associated with a specific organization ID.
// It returns a slice of pointers to Folder structs and an error, if any.
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	// Gets all folders
	folders := GetSampleData()

	// Initialises resFolder as an empty slice of folder pointers
	resFolder := []*Folder{}
	// Appends all folders with a matching orgID to resFolder
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
