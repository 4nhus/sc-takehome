# sc-grad-2025

The technical take home for SC graduate program of 2025.

## Completion notes

### Component 1

- I annotated the original ```folders.go``` file with comments explaining what I think the code does.

- I added ```folders_improvement_comments.go```, which contains the original code annotated with suggested improvements.

- I added ```folders_improved.go```, which contains the implemented improvements from above.

### Component 2

#### Token based pagination

- My solution for a paginated version of GetAllFolders implements pagination using a token.
- The token is a base64-encoded string representing the starting index of the next page.
- It throws an error if the provided token cannot be decoded or if it represents an out-of-bounds start index.
- For the same page size and token, it returns the same page (as no folders can be added to the sample data).
- Details on how it works can be gleaned from the comments and code since it is a simple implementation.

#### Why token based pagination

An alternative, simpler solution would be offset based pagination, where the FetchFolderRequest takes a starting index
for the page directly instead of a token. However, if the implementation for fetching folders used an actual database, 
then pagination by offset would be slow for a large offset value as it would unnecessarily process all the folders 
before the offset. A token for pagination could be calculated using some folder property, e.g. creation time, to filter 
through the data efficiently. 


## Getting started


Requires `Go` >= `1.20`

follow the official install instruction: [Golang Installation](https://go.dev/doc/install)

To run the code on your local machine
```
  go run main.go
```

## Folder structure

```
| go.mod
| README.md
| sample.json
| main.go
| folders
    | folders.go
    | folders_test.go
    | static.go
```

## Instructions

- This technical assessment consists of 2 components:
- Component 1:
  - within `folders.go`.
    - We would like you to read through, and run, the code.
    - Write some comments on what you think the code does.
    - Suggest some improvements that can be made to the code.
    - Implement any suggested improvements.
    - Write up some unit tests in `folders_test.go` for your new `GetAllFolders` method

- Component 2:
  - Extend your improved code to now facilitate pagination.
  - You can copy over the existing methods into `folders_pagination.go` to get started.
  - Write a short explanation of your chosen solution.

## What is pagination?
  - Pagination helps break down a large dataset into smaller chunks.
  - Those smaller chunks can then be served to the client, and are usually accompanied by a token pointing to the next chunk.
  - The end result could potentially look like this:
```
  original data: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

  This might result in the following payload to the client:
  { data: [1, 2, 3, ..., 10] }

  However, with pagination implemented, the payload might instead look like this:
  request() -> { data: [1, 2], token: "nQsjz" }

  The token could then be used to fetch more results:

  request("nQsjz") -> { data : [3, 4], token: "uJsnQ" }

  .
  .
  .

  And more results until there's no data left:

  { data: [9, 10], token: "" }
```

## Submission

Create a repo in your chosen git repository (make sure it is public so we can access it) and reply with the link to your code. We recommend using GitHub.


## Contact

If you have any questions feel free to contact us at: graduates@safetyculture.io