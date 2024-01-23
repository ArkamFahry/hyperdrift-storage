package models

import "time"

type Object struct {
	Id             string     `json:"id"`
	Name           string     `json:"name"`
	BucketId       string     `json:"bucket_id"`
	ContentType    string     `json:"content_type"`
	Size           int64      `json:"size"`
	Public         bool       `json:"public"`
	Metadata       []byte     `json:"metadata"`
	UploadStatus   string     `json:"upload_status"`
	LastAccessedAt *time.Time `json:"last_accessed_at"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type CreateObject struct {
	Id             string     `json:"id"`
	Name           string     `json:"name"`
	BucketId       string     `json:"bucket_id"`
	ContentType    string     `json:"content_type"`
	Size           int64      `json:"size"`
	Public         bool       `json:"public"`
	Metadata       []byte     `json:"metadata"`
	UploadStatus   string     `json:"upload_status"`
	LastAccessedAt *time.Time `json:"last_accessed_at"`
	CreatedAt      time.Time  `json:"created_at"`
}

type RenameObject struct {
	OldName string `json:"old_name"`
	NewName string `json:"new_name"`
}

type CopyObject struct {
	OldPath string `json:"old_path"`
	NewPath string `json:"new_path"`
}

type MoveObject struct {
	OldPath string `json:"old_path"`
	NewPath string `json:"new_path"`
}

type PreSignedObject struct {
	Url       string `json:"url"`
	Method    string `json:"method"`
	ExpiresAt int64  `json:"expires_at"`
}

type CreatePreSignedUploadObject struct {
	Bucket      string `json:"bucket"`
	Name        string `json:"name"`
	ExpiresIn   *int64 `json:"expires_in"`
	ContentType string `json:"content_type"`
	Size        int64  `json:"size"`
}

type CreatePreSignedDownloadObject struct {
	Bucket    string `json:"bucket"`
	Name      string `json:"name"`
	ExpiresIn *int64 `json:"expires_in"`
}

type CheckIfObjectExists struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

type DeleteObject struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}
