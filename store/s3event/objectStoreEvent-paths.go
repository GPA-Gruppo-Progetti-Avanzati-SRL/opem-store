package s3event

// This file contains the paths to the field in the generated entity.
// A path is a string with all the identifiers from the root document to the single leaves.
// In case of maps and arrays place holder for the key (%s) or the index %d have been provided.

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

const (
	OIdFieldName         = "_id"
	BidFieldName         = "_bid"
	EtFieldName          = "_et"
	RipFieldName         = "_rip"
	StatusFieldName      = "_status"
	Status_CodeFieldName = "_status.code"
	Status_TextFieldName = "_status.text"
	EventFieldName       = "event"
)

// @tpm-schematics:start-region("bottom-file-section")

const (
	EventS3BucketNameFieldName = "event.s3.bucket.name"
)

// @tpm-schematics:end-region("bottom-file-section")
