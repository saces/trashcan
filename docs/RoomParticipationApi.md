# \RoomParticipationApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DefineFilter**](RoomParticipationApi.md#DefineFilter) | **Post** /_matrix/client/unstable/user/{userId}/filter | Upload a new filter.
[**GetEventContext**](RoomParticipationApi.md#GetEventContext) | **Get** /_matrix/client/unstable/rooms/{roomId}/context/{eventId} | Get events and state around the specified event.
[**GetEvents**](RoomParticipationApi.md#GetEvents) | **Get** /_matrix/client/unstable/events | Listen on the event stream.
[**GetFilter**](RoomParticipationApi.md#GetFilter) | **Get** /_matrix/client/unstable/user/{userId}/filter/{filterId} | Download a filter
[**GetJoinedMembersByRoom**](RoomParticipationApi.md#GetJoinedMembersByRoom) | **Get** /_matrix/client/unstable/rooms/{roomId}/joined_members | Gets the list of currently joined users and their profile data.
[**GetMembersByRoom**](RoomParticipationApi.md#GetMembersByRoom) | **Get** /_matrix/client/unstable/rooms/{roomId}/members | Get the m.room.member events for the room.
[**GetOneEvent**](RoomParticipationApi.md#GetOneEvent) | **Get** /_matrix/client/unstable/events/{eventId} | Get a single event by event ID.
[**GetOneRoomEvent**](RoomParticipationApi.md#GetOneRoomEvent) | **Get** /_matrix/client/unstable/rooms/{roomId}/event/{eventId} | Get a single event by event ID.
[**GetRoomEvents**](RoomParticipationApi.md#GetRoomEvents) | **Get** /_matrix/client/unstable/rooms/{roomId}/messages | Get a list of events for this room
[**GetRoomState**](RoomParticipationApi.md#GetRoomState) | **Get** /_matrix/client/unstable/rooms/{roomId}/state | Get all state events in the current state of a room.
[**GetRoomStateWithKey**](RoomParticipationApi.md#GetRoomStateWithKey) | **Get** /_matrix/client/unstable/rooms/{roomId}/state/{eventType}/{stateKey} | Get the state identified by the type and key.
[**InitialSync**](RoomParticipationApi.md#InitialSync) | **Get** /_matrix/client/unstable/initialSync | Get the user&#39;s current state.
[**PostReceipt**](RoomParticipationApi.md#PostReceipt) | **Post** /_matrix/client/unstable/rooms/{roomId}/receipt/{receiptType}/{eventId} | Send a receipt for the given event ID.
[**RedactEvent**](RoomParticipationApi.md#RedactEvent) | **Put** /_matrix/client/unstable/rooms/{roomId}/redact/{eventId}/{txnId} | Strips all non-integrity-critical information out of an event.
[**RoomInitialSync**](RoomParticipationApi.md#RoomInitialSync) | **Get** /_matrix/client/unstable/rooms/{roomId}/initialSync | Snapshot the current state of a room and its most recent messages.
[**SendMessage**](RoomParticipationApi.md#SendMessage) | **Put** /_matrix/client/unstable/rooms/{roomId}/send/{eventType}/{txnId} | Send a message event to the given room.
[**SetRoomStateWithKey**](RoomParticipationApi.md#SetRoomStateWithKey) | **Put** /_matrix/client/unstable/rooms/{roomId}/state/{eventType}/{stateKey} | Send a state event to the given room.
[**SetTyping**](RoomParticipationApi.md#SetTyping) | **Put** /_matrix/client/unstable/rooms/{roomId}/typing/{userId} | Informs the server that the user has started or stopped typing.
[**Sync**](RoomParticipationApi.md#Sync) | **Get** /_matrix/client/unstable/sync | Synchronise the client&#39;s state and receive new messages.



## DefineFilter

> InlineResponse20040 DefineFilter(ctx, userId, filter)
Upload a new filter.

Uploads a new filter definition to the homeserver. Returns a filter ID that may be used in future requests to restrict which events are returned to the client.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The id of the user uploading the filter. The access token must be authorized to make requests for this user id. | 
**filter** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| The filter to upload. | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventContext

> InlineResponse20033 GetEventContext(ctx, roomId, eventId, optional)
Get events and state around the specified event.

This API returns a number of events that happened just before and after the specified event. This allows clients to get the context surrounding an event.  *Note*: This endpoint supports lazy-loading of room member events. See `Lazy-loading room members <#lazy-loading-room-members>`_ for more information.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room to get events from. | 
**eventId** | **string**| The event to get context around. | 
 **optional** | ***GetEventContextOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventContextOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| The maximum number of events to return. Default: 10. | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvents

> InlineResponse20010 GetEvents(ctx, optional)
Listen on the event stream.

This will listen for new events and return them to the caller. This will block until an event is received, or until the ``timeout`` is reached.  This endpoint was deprecated in r0 of this specification. Clients should instead call the |/sync|_ API with a ``since`` parameter. See the `migration guide <https://matrix.org/docs/guides/client-server-migrating-from-v1.html#deprecated-endpoints>`_.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **optional.String**| The token to stream from. This token is either from a previous request to this API or from the initial sync API. | 
 **timeout** | **optional.Int32**| The maximum time in milliseconds to wait for an event. | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilter

> InlineResponse20041 GetFilter(ctx, userId, filterId)
Download a filter

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The user ID to download a filter for. | 
**filterId** | **string**| The filter ID to download. | 

### Return type

[**InlineResponse20041**](inline_response_200_41.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJoinedMembersByRoom

> InlineResponse20034 GetJoinedMembersByRoom(ctx, roomId)
Gets the list of currently joined users and their profile data.

This API returns a map of MXIDs to member info objects for members of the room. The current user must be in the room for it to work, unless it is an Application Service in which case any of the AS's users must be in the room. This API is primarily for Application Services and should be faster to respond than ``/members`` as it can be implemented more efficiently on the server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room to get the members of. | 

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMembersByRoom

> InlineResponse20035 GetMembersByRoom(ctx, roomId, optional)
Get the m.room.member events for the room.

Get the list of members for this room.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room to get the member events for. | 
 **optional** | ***GetMembersByRoomOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMembersByRoomOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **at** | **optional.String**| The point in time (pagination token) to return members for in the room. This token can be obtained from a &#x60;&#x60;prev_batch&#x60;&#x60; token returned for each room by the sync API. Defaults to the current state of the room, as determined by the server. | 
 **membership** | **optional.String**| The kind of membership to filter for. Defaults to no filtering if unspecified. When specified alongside &#x60;&#x60;not_membership&#x60;&#x60;, the two parameters create an &#39;or&#39; condition: either the membership *is* the same as &#x60;&#x60;membership&#x60;&#x60; **or** *is not* the same as &#x60;&#x60;not_membership&#x60;&#x60;. | 
 **notMembership** | **optional.String**| The kind of membership to exclude from the results. Defaults to no filtering if unspecified. | 

### Return type

[**InlineResponse20035**](inline_response_200_35.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOneEvent

> InlineResponse20011 GetOneEvent(ctx, eventId)
Get a single event by event ID.

Get a single event based on ``event_id``. You must have permission to retrieve this event e.g. by being a member in the room for this event.  This endpoint was deprecated in r0 of this specification. Clients should instead call the |/rooms/{roomId}/event/{eventId}|_ API or the |/rooms/{roomId}/context/{eventId}|_ API.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| The event ID to get. | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOneRoomEvent

> InlineResponse20011 GetOneRoomEvent(ctx, roomId, eventId)
Get a single event by event ID.

Get a single event based on ``roomId/eventId``. You must have permission to retrieve this event e.g. by being a member in the room for this event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The ID of the room the event is in. | 
**eventId** | **string**| The event ID to get. | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomEvents

> InlineResponse20036 GetRoomEvents(ctx, roomId, from, dir, optional)
Get a list of events for this room

This API returns a list of message and state events for a room. It uses pagination query parameters to paginate history in the room.  *Note*: This endpoint supports lazy-loading of room member events. See `Lazy-loading room members <#lazy-loading-room-members>`_ for more information.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room to get events from. | 
**from** | **string**| The token to start returning events from. This token can be obtained from a &#x60;&#x60;prev_batch&#x60;&#x60; token returned for each room by the sync API, or from a &#x60;&#x60;start&#x60;&#x60; or &#x60;&#x60;end&#x60;&#x60; token returned by a previous request to this endpoint. | 
**dir** | **string**| The direction to return events from. | 
 **optional** | ***GetRoomEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRoomEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **to** | **optional.String**| The token to stop returning events at. This token can be obtained from a &#x60;&#x60;prev_batch&#x60;&#x60; token returned for each room by the sync endpoint, or from a &#x60;&#x60;start&#x60;&#x60; or &#x60;&#x60;end&#x60;&#x60; token returned by a previous request to this endpoint. | 
 **limit** | **optional.Int32**| The maximum number of events to return. Default: 10. | 
 **filter** | **optional.String**| A JSON RoomEventFilter to filter returned events with. | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomState

> []map[string]interface{} GetRoomState(ctx, roomId)
Get all state events in the current state of a room.

Get the state events for the current state of a room.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room to look up the state for. | 

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomStateWithKey

> map[string]interface{} GetRoomStateWithKey(ctx, roomId, eventType, stateKey)
Get the state identified by the type and key.

.. For backwards compatibility with older links... .. _`get-matrix-client-unstable-rooms-roomid-state-eventtype`:  Looks up the contents of a state event in a room. If the user is joined to the room then the state is taken from the current state of the room. If the user has left the room then the state is taken from the state of the room when they left.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room to look up the state in. | 
**eventType** | **string**| The type of state to look up. | 
**stateKey** | **string**| The key of the state to look up. Defaults to an empty string. When an empty string, the trailing slash on this endpoint is optional. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitialSync

> InlineResponse20012 InitialSync(ctx, optional)
Get the user's current state.

This returns the full state for this user, with an optional limit on the number of messages per room to return.  This endpoint was deprecated in r0 of this specification. Clients should instead call the |/sync|_ API with no ``since`` parameter. See the `migration guide <https://matrix.org/docs/guides/client-server-migrating-from-v1.html#deprecated-endpoints>`_.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InitialSyncOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InitialSyncOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The maximum number of messages to return for each room. | 
 **archived** | **optional.Bool**| Whether to include rooms that the user has left. If &#x60;&#x60;false&#x60;&#x60; then only rooms that the user has been invited to or has joined are included. If set to &#x60;&#x60;true&#x60;&#x60; then rooms that the user has left are included as well. By default this is &#x60;&#x60;false&#x60;&#x60;. | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostReceipt

> map[string]interface{} PostReceipt(ctx, roomId, receiptType, eventId, optional)
Send a receipt for the given event ID.

This API updates the marker for the given receipt type to the event ID specified.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room in which to send the event. | 
**receiptType** | **string**| The type of receipt to send. | 
**eventId** | **string**| The event ID to acknowledge up to. | 
 **optional** | ***PostReceiptOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostReceiptOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **receipt** | **optional.Map[string]interface{}**| Extra receipt information to attach to &#x60;&#x60;content&#x60;&#x60; if any. The server will automatically set the &#x60;&#x60;ts&#x60;&#x60; field. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedactEvent

> InlineResponse20037 RedactEvent(ctx, roomId, eventId, txnId, optional)
Strips all non-integrity-critical information out of an event.

Strips all information out of an event which isn't critical to the integrity of the server-side representation of the room.  This cannot be undone.  Users may redact their own events, and any user with a power level greater than or equal to the `redact` power level of the room may redact events there.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room from which to redact the event. | 
**eventId** | **string**| The ID of the event to redact | 
**txnId** | **string**| The transaction ID for this event. Clients should generate a unique ID; it will be used by the server to ensure idempotency of requests. | 
 **optional** | ***RedactEventOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RedactEventOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of InlineObject30**](InlineObject30.md)|  | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoomInitialSync

> RoomInfo1 RoomInitialSync(ctx, roomId)
Snapshot the current state of a room and its most recent messages.

Get a copy of the current state and the most recent messages in a room.  This endpoint was deprecated in r0 of this specification. There is no direct replacement; the relevant information is returned by the |/sync|_ API. See the `migration guide <https://matrix.org/docs/guides/client-server-migrating-from-v1.html#deprecated-endpoints>`_.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room to get the data. | 

### Return type

[**RoomInfo1**](RoomInfo_1.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMessage

> InlineResponse20037 SendMessage(ctx, roomId, eventType, txnId, optional)
Send a message event to the given room.

This endpoint is used to send a message event to a room. Message events allow access to historical events and pagination, making them suited for \"once-off\" activity in a room.  The body of the request should be the content object of the event; the fields in this object will vary depending on the type of event. See `Room Events`_ for the m. event specification.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room to send the event to. | 
**eventType** | **string**| The type of event to send. | 
**txnId** | **string**| The transaction ID for this event. Clients should generate an ID unique across requests with the same access token; it will be used by the server to ensure idempotency of requests. | 
 **optional** | ***SendMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SendMessageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **optional.Map[string]interface{}**|  | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRoomStateWithKey

> InlineResponse20037 SetRoomStateWithKey(ctx, roomId, eventType, stateKey, optional)
Send a state event to the given room.

.. For backwards compatibility with older links... .. _`put-matrix-client-unstable-rooms-roomid-state-eventtype`:  State events can be sent using this endpoint.  These events will be overwritten if ``<room id>``, ``<event type>`` and ``<state key>`` all match.  Requests to this endpoint **cannot use transaction IDs** like other ``PUT`` paths because they cannot be differentiated from the ``state_key``. Furthermore, ``POST`` is unsupported on state paths.  The body of the request should be the content object of the event; the fields in this object will vary depending on the type of event. See `Room Events`_ for the ``m.`` event specification. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room to set the state in | 
**eventType** | **string**| The type of event to send. | 
**stateKey** | **string**| The state_key for the state to send. Defaults to the empty string. When an empty string, the trailing slash on this endpoint is optional. | 
 **optional** | ***SetRoomStateWithKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetRoomStateWithKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **optional.Map[string]interface{}**|  | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTyping

> map[string]interface{} SetTyping(ctx, userId, roomId, typingState)
Informs the server that the user has started or stopped typing.

This tells the server that the user is typing for the next N milliseconds where N is the value specified in the ``timeout`` key. Alternatively, if ``typing`` is ``false``, it tells the server that the user has stopped typing.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The user who has started to type. | 
**roomId** | **string**| The room in which the user is typing. | 
**typingState** | [**InlineObject32**](InlineObject32.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Sync

> InlineResponse20039 Sync(ctx, optional)
Synchronise the client's state and receive new messages.

Synchronise the client's state with the latest state on the server. Clients use this API when they first log in to get an initial snapshot of the state on the server, and then continue to call this API to get incremental deltas to the state, and to receive new messages.  *Note*: This endpoint supports lazy-loading. See `Filtering <#filtering>`_ for more information. Lazy-loading members is only supported on a ``StateFilter`` for this endpoint. When lazy-loading is enabled, servers MUST include the syncing user's own membership event when they join a room, or when the full state of rooms is requested, to aid discovering the user's avatar & displayname.  Like other members, the user's own membership event is eligible for being considered redundant by the server. When a sync is ``limited``, the server MUST return membership events for events in the gap (between ``since`` and the start of the returned timeline), regardless as to whether or not they are redundant.  This ensures that joins/leaves and profile changes which occur during the gap are not lost.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SyncOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SyncOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| The ID of a filter created using the filter API or a filter JSON object encoded as a string. The server will detect whether it is an ID or a JSON object by whether the first character is a &#x60;&#x60;\&quot;{\&quot;&#x60;&#x60; open brace. Passing the JSON inline is best suited to one off requests. Creating a filter using the filter API is recommended for clients that reuse the same filter multiple times, for example in long poll requests.  See &#x60;Filtering &lt;#filtering&gt;&#x60;_ for more information. | 
 **since** | **optional.String**| A point in time to continue a sync from. | 
 **fullState** | **optional.Bool**| Controls whether to include the full state for all rooms the user is a member of.  If this is set to &#x60;&#x60;true&#x60;&#x60;, then all state events will be returned, even if &#x60;&#x60;since&#x60;&#x60; is non-empty. The timeline will still be limited by the &#x60;&#x60;since&#x60;&#x60; parameter. In this case, the &#x60;&#x60;timeout&#x60;&#x60; parameter will be ignored and the query will return immediately, possibly with an empty timeline.  If &#x60;&#x60;false&#x60;&#x60;, and &#x60;&#x60;since&#x60;&#x60; is non-empty, only state which has changed since the point indicated by &#x60;&#x60;since&#x60;&#x60; will be returned.  By default, this is &#x60;&#x60;false&#x60;&#x60;. | 
 **setPresence** | **optional.String**| Controls whether the client is automatically marked as online by polling this API. If this parameter is omitted then the client is automatically marked as online when it uses this API. Otherwise if the parameter is set to \&quot;offline\&quot; then the client is not marked as being online when it uses this API. When set to \&quot;unavailable\&quot;, the client is marked as being idle. | 
 **timeout** | **optional.Int32**| The maximum time to wait, in milliseconds, before returning this request. If no events (or other data) become available before this time elapses, the server will return a response with empty fields.  By default, this is &#x60;&#x60;0&#x60;&#x60;, so the server will return immediately even if the response is empty. | 

### Return type

[**InlineResponse20039**](inline_response_200_39.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

