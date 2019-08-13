# InlineResponse20031

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | An access token for the account. This access token can then be used to authorize other requests. Required if the &#x60;&#x60;inhibit_login&#x60;&#x60; option is false. | [optional] 
**DeviceId** | **string** | ID of the registered device. Will be the same as the corresponding parameter in the request, if one was specified. Required if the &#x60;&#x60;inhibit_login&#x60;&#x60; option is false. | [optional] 
**HomeServer** | **string** | The server_name of the homeserver on which the account has been registered.  **Deprecated**. Clients should extract the server_name from &#x60;&#x60;user_id&#x60;&#x60; (by splitting at the first colon) if they require it. Note also that &#x60;&#x60;homeserver&#x60;&#x60; is not spelt this way. | [optional] 
**UserId** | **string** | The fully-qualified Matrix user ID (MXID) that has been registered.  Any user ID returned by this API must conform to the grammar given in the &#x60;Matrix specification &lt;../appendices.html#user-identifiers&gt;&#x60;_. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


