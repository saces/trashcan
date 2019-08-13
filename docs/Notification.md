# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | **[]string** | The action(s) to perform when the conditions for this rule are met. See &#x60;Push Rules: API&#x60;_. | 
**Event** | [**map[string]interface{}**](map[string]interface{}.md) | The Event object for the event that triggered the notification. | 
**ProfileTag** | **string** | The profile tag of the rule that matched this event. | [optional] 
**Read** | **bool** | Indicates whether the user has sent a read receipt indicating that they have read this message. | 
**RoomId** | **string** | The ID of the room in which the event was posted. | 
**Ts** | **int32** | The unix timestamp at which the event notification was sent, in milliseconds. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


