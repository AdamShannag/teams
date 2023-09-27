package message

const (
	CREATED_SUCCESSFULLY        = "Successfully created a new %s, uuid: [%s]"
	DELETED_SUCCESSFULLY        = "Successfully deleted a list of  %s, uuid's: %s"
	UPDATED_SUCCESSFULLY        = "Successfully updated %s, uuid: [%s]"
	RETRIEVED_SUCCESSFULLY      = "Successfully retrieved %s, uuid: [%s]"
	RETRIEVED_LIST_SUCCESSFULLY = "Successfully retrieved a list of %s, size: [%d]"

	CREATED_FAILED        = "Failed to create a new %s"
	DELETED_FAILED        = "Failed to delete the list of %s, uuid's: %s"
	UPDATED_FAILED        = "Failed to update %s, uuid: [%s]"
	RETRIEVED_FAILED      = "Failed to retrieve %s, uuid: [%s]"
	RETRIEVED_LIST_FAILED = "Failed to retrieve the list of %s"

	FAILED_PARSING = "Failed to parse the request JSON"
)
