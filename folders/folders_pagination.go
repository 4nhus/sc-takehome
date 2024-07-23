package folders

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

// GetAllFoldersWithOrgIDPaginated retrieves a token-based page of folders for a given orgID. It takes a
// FetchFolderRequest, which includes the orgID, page size, and a token representing the page start index. The function
// returns a FetchFolderResponse containing a slice of folder pointers and a token for the next page if it exists,
// else an empty string. If the token is cannot be decoded or represents an out-of-bounds page start index, an error
// will be returned.
func GetAllFoldersWithOrgIDPaginated(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// Decode the token to get the page start index, with an early return for errors
	start := 0
	if req.Token != "" {
		decodedToken, err := base64.StdEncoding.DecodeString(req.Token)
		if err != nil {
			return nil, fmt.Errorf("invalid token: cannot be decoded")
		}
		start, err = strconv.Atoi(string(decodedToken))
		if err != nil {
			return nil, fmt.Errorf("invalid token: cannot be decoded")
		}
	}

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

	numFolders := len(foldersWithOrgID)
	end := start + req.PageSize

	// Throw an error if the page start index is out of bounds
	if start < 0 || start >= numFolders {
		return nil, fmt.Errorf("invalid token: out of bounds")
	}
	// Clamp end index to keep it in bounds
	if end > numFolders {
		end = numFolders
	}

	paginatedFoldersWithOrgId := foldersWithOrgID[start:end]

	// Create the next token
	nextToken := ""
	if end < numFolders {
		nextToken = base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(end)))
	}

	return &FetchFolderResponse{Folders: paginatedFoldersWithOrgId, Token: nextToken}, nil
}
