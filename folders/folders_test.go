package folders_test

import (
	"errors"
	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	//"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	// Save the original function
	originalFetchAllFoldersFunc := folders.FetchAllFoldersFunc
	// Restore the original function after the test
	defer func() { folders.FetchAllFoldersFunc = originalFetchAllFoldersFunc }()

	orgID1 := uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111")
	orgID2 := uuid.FromStringOrNil("22222222-2222-2222-2222-222222222222")
	fetchRequest := &folders.FetchFolderRequest{OrgID: orgID1}

	folder1 := &folders.Folder{OrgId: orgID1}
	folder2 := &folders.Folder{OrgId: orgID1}
	folder3 := &folders.Folder{OrgId: orgID2}
	folder4 := &folders.Folder{OrgId: orgID2}

	mockFolders1 := []*folders.Folder{
		folder1,
		folder2,
	}

	mockFolders2 := []*folders.Folder{
		folder1,
		folder2,
		folder3,
		folder4,
	}

	mockFolders3 := []*folders.Folder{
		folder3,
		folder4,
	}

	var emptyFolders []*folders.Folder

	// Could parameterise these test cases
	t.Run("All folders match orgID", func(t *testing.T) {
		folders.FetchAllFoldersFunc = func() ([]*folders.Folder, error) {
			return mockFolders1, nil
		}

		res, err := folders.GetAllFoldersWithOrgID(fetchRequest)
		expectedFolders := []*folders.Folder{
			folder1,
			folder2,
		}

		assert.Equal(t, expectedFolders, res.Folders)
		assert.Nil(t, err)
	})

	t.Run("Some folders match orgID", func(t *testing.T) {
		folders.FetchAllFoldersFunc = func() ([]*folders.Folder, error) {
			return mockFolders2, nil
		}

		res, err := folders.GetAllFoldersWithOrgID(fetchRequest)
		expectedFolders := []*folders.Folder{
			folder1,
			folder2,
		}

		assert.Equal(t, expectedFolders, res.Folders)
		assert.Nil(t, err)
	})

	t.Run("No folders match orgID", func(t *testing.T) {
		folders.FetchAllFoldersFunc = func() ([]*folders.Folder, error) {
			return mockFolders3, nil
		}

		res, err := folders.GetAllFoldersWithOrgID(fetchRequest)
		expectedFolders := emptyFolders

		assert.Equal(t, expectedFolders, res.Folders)
		assert.Nil(t, err)
	})

	t.Run("No folders", func(t *testing.T) {
		folders.FetchAllFoldersFunc = func() ([]*folders.Folder, error) {
			return emptyFolders, nil
		}

		res, err := folders.GetAllFoldersWithOrgID(fetchRequest)
		expectedFolders := emptyFolders

		assert.Equal(t, expectedFolders, res.Folders)
		assert.Nil(t, err)
	})

	t.Run("Error fetching folders", func(t *testing.T) {
		folders.FetchAllFoldersFunc = func() ([]*folders.Folder, error) {
			return nil, errors.New("error")
		}

		res, err := folders.GetAllFoldersWithOrgID(fetchRequest)

		assert.Nil(t, res)
		assert.Error(t, err)
		assert.Equal(t, "error", err.Error())
	})
}
