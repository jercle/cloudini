package m365

type AzureGraphResponse[T MailboxUsageDetail | interface{}] struct {
	Context  *string `json:"@odata.context" bson:"@odata.context"`
	NextLink *string `json:"@odata.nextLink" bson:"@odata.nextLink"`
	Value    []T     `json:"value" bson:"value"`
}
