# InlineObject23

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | [**map[string]map[string]interface{}**](map[string]interface{}.md) | Additional authentication information for the user-interactive authentication API. Note that this information is *not* used to define how the registered user should be authenticated, but is instead used to authenticate the &#x60;&#x60;register&#x60;&#x60; call itself. | [optional] 
**BindEmail** | **bool** | If true, the server binds the email used for authentication to the Matrix ID with the identity server. | [optional] 
**BindMsisdn** | **bool** | If true, the server binds the phone number used for authentication to the Matrix ID with the identity server. | [optional] 
**DeviceId** | **string** | ID of the client device. If this does not correspond to a known client device, a new device will be created. The server will auto-generate a device_id if this is not specified. | [optional] 
**InhibitLogin** | **bool** | If true, an &#x60;&#x60;access_token&#x60;&#x60; and &#x60;&#x60;device_id&#x60;&#x60; should not be returned from this call, therefore preventing an automatic login. Defaults to false. | [optional] 
**InitialDeviceDisplayName** | **string** | A display name to assign to the newly-created device. Ignored if &#x60;&#x60;device_id&#x60;&#x60; corresponds to a known device. | [optional] 
**Password** | **string** | The desired password for the account. | [optional] 
**Username** | **string** | The basis for the localpart of the desired Matrix ID. If omitted, the homeserver MUST generate a Matrix ID local part. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


