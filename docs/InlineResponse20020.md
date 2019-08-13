# InlineResponse20020

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | An access token for the account. This access token can then be used to authorize other requests. | [optional] 
**DeviceId** | **string** | ID of the logged-in device. Will be the same as the corresponding parameter in the request, if one was specified. | [optional] 
**HomeServer** | **string** | The server_name of the homeserver on which the account has been registered.  **Deprecated**. Clients should extract the server_name from &#x60;&#x60;user_id&#x60;&#x60; (by splitting at the first colon) if they require it. Note also that &#x60;&#x60;homeserver&#x60;&#x60; is not spelt this way. | [optional] 
**UserId** | **string** | The fully-qualified Matrix ID that has been registered. | [optional] 
**WellKnown** | [**map[string]map[string]interface{}**](map[string]interface{}.md) | Optional client configuration provided by the server. If present, clients SHOULD use the provided object to reconfigure themselves, optionally validating the URLs within. This object takes the same form as the one returned from .well-known autodiscovery. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


