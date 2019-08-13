# AuthenticationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | **[]string** | A list of the stages the client has completed successfully | [optional] 
**Flows** | [**[]AuthenticationResponseFlows**](Authentication_response_flows.md) | A list of the login flows supported by the server for this API. | 
**Params** | [**map[string]map[string]interface{}**](map[string]interface{}.md) | Contains any information that the client will need to know in order to use a given type of authentication. For each login type presented, that type may be present as a key in this dictionary. For example, the public part of an OAuth client ID could be given here. | [optional] 
**Session** | **string** | This is a session identifier that the client must pass back to the home server, if one is provided, in subsequent attempts to authenticate in the same API call. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


