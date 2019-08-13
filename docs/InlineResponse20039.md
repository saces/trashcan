# InlineResponse20039

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountData** | [**map[string]interface{}**](map[string]interface{}.md) | The global private data created by this user. | [optional] 
**DeviceLists** | [**map[string]interface{}**](.md) | Information on end-to-end device updates, as specified in |device_lists_sync|_. | [optional] 
**DeviceOneTimeKeysCount** | **map[string]int32** | Information on end-to-end encryption keys, as specified in |device_lists_sync|_. | [optional] 
**NextBatch** | **string** | The batch token to supply in the &#x60;&#x60;since&#x60;&#x60; param of the next &#x60;&#x60;/sync&#x60;&#x60; request. | 
**Presence** | [**map[string]interface{}**](map[string]interface{}.md) | The updates to the presence status of other users. | [optional] 
**Rooms** | [**Rooms**](Rooms.md) |  | [optional] 
**ToDevice** | [**map[string]interface{}**](.md) | Information on the send-to-device messages for the client device, as defined in |send_to_device_sync|_. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


