package dns

//Modify with your Access Key Id and Access Key Secret
const (
//	TestAccessKeyId     = "MY_ACCESS_KEY_ID"
//	TestAccessKeySecret = "MY_ACCESS_KEY_SECRET"
	TestAccessKeyId     = "1RYsYarqhhWKQnqn"
	TestAccessKeySecret = "mppAp28laOPBRQ97nByZtpdXU5zerR"
	TestDomainName = "aisafe.win"
)

var testClient *Client

func NewTestClient() *Client {
	if testClient == nil {
		testClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
	}
	return testClient
}

var testDebugClient *Client

func NewTestClientForDebug() *Client {
	if testDebugClient == nil {
		testDebugClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
		testDebugClient.SetDebug(true)
	}
	return testDebugClient
}
