# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/runtime/v1/registry.proto](#atomix_runtime_v1_registry-proto)
    - [DriverInfo](#atomix-runtime-v1-DriverInfo)
    - [PluginChunk](#atomix-runtime-v1-PluginChunk)
    - [PluginHeader](#atomix-runtime-v1-PluginHeader)
    - [PluginTrailer](#atomix-runtime-v1-PluginTrailer)
    - [PullDriverRequest](#atomix-runtime-v1-PullDriverRequest)
    - [PullDriverResponse](#atomix-runtime-v1-PullDriverResponse)
    - [PushDriverRequest](#atomix-runtime-v1-PushDriverRequest)
    - [PushDriverResponse](#atomix-runtime-v1-PushDriverResponse)
    - [RuntimeInfo](#atomix-runtime-v1-RuntimeInfo)
  
    - [Registry](#atomix-runtime-v1-Registry)
  
- [Scalar Value Types](#scalar-value-types)



<a name="atomix_runtime_v1_registry-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/runtime/v1/registry.proto



<a name="atomix-runtime-v1-DriverInfo"></a>

### DriverInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| version | [string](#string) |  |  |






<a name="atomix-runtime-v1-PluginChunk"></a>

### PluginChunk



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [bytes](#bytes) |  |  |






<a name="atomix-runtime-v1-PluginHeader"></a>

### PluginHeader



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| driver | [DriverInfo](#atomix-runtime-v1-DriverInfo) |  |  |
| runtime | [RuntimeInfo](#atomix-runtime-v1-RuntimeInfo) |  |  |






<a name="atomix-runtime-v1-PluginTrailer"></a>

### PluginTrailer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| checksum | [string](#string) |  |  |






<a name="atomix-runtime-v1-PullDriverRequest"></a>

### PullDriverRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [PluginHeader](#atomix-runtime-v1-PluginHeader) |  |  |






<a name="atomix-runtime-v1-PullDriverResponse"></a>

### PullDriverResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| chunk | [PluginChunk](#atomix-runtime-v1-PluginChunk) |  |  |
| trailer | [PluginTrailer](#atomix-runtime-v1-PluginTrailer) |  |  |






<a name="atomix-runtime-v1-PushDriverRequest"></a>

### PushDriverRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [PluginHeader](#atomix-runtime-v1-PluginHeader) |  |  |
| chunk | [PluginChunk](#atomix-runtime-v1-PluginChunk) |  |  |
| trailer | [PluginTrailer](#atomix-runtime-v1-PluginTrailer) |  |  |






<a name="atomix-runtime-v1-PushDriverResponse"></a>

### PushDriverResponse







<a name="atomix-runtime-v1-RuntimeInfo"></a>

### RuntimeInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  |  |





 

 

 


<a name="atomix-runtime-v1-Registry"></a>

### Registry
The registry service provides control functions for the runtime.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| PushDriver | [PushDriverRequest](#atomix-runtime-v1-PushDriverRequest) stream | [PushDriverResponse](#atomix-runtime-v1-PushDriverResponse) |  |
| PullDriver | [PullDriverRequest](#atomix-runtime-v1-PullDriverRequest) | [PullDriverResponse](#atomix-runtime-v1-PullDriverResponse) stream |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

