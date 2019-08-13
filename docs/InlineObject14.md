# InlineObject14

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Third party identifier for the user.  Deprecated in favour of &#x60;&#x60;identifier&#x60;&#x60;. | [optional] 
**DeviceId** | **string** | ID of the client device. If this does not correspond to a known client device, a new device will be created. The server will auto-generate a device_id if this is not specified. | [optional] 
**Identifier** | [**UserIdentifier**](User identifier.md) |  | [optional] 
**InitialDeviceDisplayName** | **string** | A display name to assign to the newly-created device. Ignored if &#x60;&#x60;device_id&#x60;&#x60; corresponds to a known device. | [optional] 
**Medium** | **string** | When logging in using a third party identifier, the medium of the identifier. Must be &#39;email&#39;.  Deprecated in favour of &#x60;&#x60;identifier&#x60;&#x60;. | [optional] 
**Password** | **string** | Required when &#x60;&#x60;type&#x60;&#x60; is &#x60;&#x60;m.login.password&#x60;&#x60;. The user&#39;s password. | [optional] 
**Token** | **string** | Required when &#x60;&#x60;type&#x60;&#x60; is &#x60;&#x60;m.login.token&#x60;&#x60;. Part of &#x60;Token-based&#x60;_ login. | [optional] 
**Type** | **string** | The login type being used. | 
**User** | **string** | The fully qualified user ID or just local part of the user ID, to log in.  Deprecated in favour of &#x60;&#x60;identifier&#x60;&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


