package log

const (
	CREATED_SUCCESSFULLY        = "Successfully created a new %s, uuid: [%s]"
	DELETED_SUCCESSFULLY        = "Successfully deleted a List of  %s, uuid's: %s"
	UPDATED_SUCCESSFULLY        = "Successfully updated %s, uuid: [%s]"
	RETRIEVED_SUCCESSFULLY      = "Successfully retrieved %s, uuid: [%s]"
	RETRIEVED_LIST_SUCCESSFULLY = "Successfully retrieved a List of %s, size: [%d]"

	CREATED_FAILED        = "Failed to Create a new %s"
	DELETED_FAILED        = "Failed to Delete the List of %s, uuid's: %s"
	UPDATED_FAILED        = "Failed to Update %s, uuid: [%s]"
	RETRIEVED_FAILED      = "Failed to Retrieve %s, uuid: [%s]"
	RETRIEVED_LIST_FAILED = "Failed to Retrieve the List of %s"

	FAILED_PARSING = "Failed to parse the request JSON"
)
