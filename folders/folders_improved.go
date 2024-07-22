package folders

// For mocking in unit tests
var FetchAllFoldersFunc = FetchAllFolders

// GetAllFoldersWithOrgID retrieves all folders for a given orgID.
// It takes a FetchFolderRequest - which wraps an orgID - and returns a FetchFolderResponse - which wraps a slice of
// folder pointers - and an error, if any.
func GetAllFoldersWithOrgID(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// FetchAllFolders may generate errors if it is replaced with an actual network / database implementation.
	folders, err := FetchAllFoldersFunc()
	if err != nil {
		return nil, err
	}

	var foldersWithOrgID []*Folder
	for _, folder := range folders {
		if folder.OrgId == req.OrgID {
			foldersWithOrgID = append(foldersWithOrgID, folder)
		}
	}

	return &FetchFolderResponse{Folders: foldersWithOrgID}, nil
}

// FetchAllFolders fetches all folders, and returns them and an error if any.
func FetchAllFolders() ([]*Folder, error) {
	return GetSampleData(), nil
}
