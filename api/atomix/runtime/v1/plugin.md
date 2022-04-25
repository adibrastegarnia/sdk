# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/runtime/v1/plugin.proto](#atomix_runtime_v1_plugin-proto)
    - [CompilePluginRequest](#atomix-runtime-v1-CompilePluginRequest)
    - [CompilePluginResponse](#atomix-runtime-v1-CompilePluginResponse)
    - [PluginInfo](#atomix-runtime-v1-PluginInfo)
    - [PullPluginRequest](#atomix-runtime-v1-PullPluginRequest)
    - [PullPluginResponse](#atomix-runtime-v1-PullPluginResponse)
    - [RuntimeInfo](#atomix-runtime-v1-RuntimeInfo)
  
    - [PluginCompiler](#atomix-runtime-v1-PluginCompiler)
    - [PluginRegistry](#atomix-runtime-v1-PluginRegistry)
  
- [Scalar Value Types](#scalar-value-types)



<a name="atomix_runtime_v1_plugin-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/runtime/v1/plugin.proto



<a name="atomix-runtime-v1-CompilePluginRequest"></a>

### CompilePluginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| runtime | [RuntimeInfo](#atomix-runtime-v1-RuntimeInfo) |  |  |
| plugin | [PluginInfo](#atomix-runtime-v1-PluginInfo) |  |  |






<a name="atomix-runtime-v1-CompilePluginResponse"></a>

### CompilePluginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| plugin | [string](#string) |  |  |






<a name="atomix-runtime-v1-PluginInfo"></a>

### PluginInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| version | [string](#string) |  |  |






<a name="atomix-runtime-v1-PullPluginRequest"></a>

### PullPluginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| runtime | [RuntimeInfo](#atomix-runtime-v1-RuntimeInfo) |  |  |
| plugin | [PluginInfo](#atomix-runtime-v1-PluginInfo) |  |  |






<a name="atomix-runtime-v1-PullPluginResponse"></a>

### PullPluginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [bytes](#bytes) |  |  |






<a name="atomix-runtime-v1-RuntimeInfo"></a>

### RuntimeInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  |  |
| build_version | [string](#string) |  |  |





 

 

 


<a name="atomix-runtime-v1-PluginCompiler"></a>

### PluginCompiler
The plugin compiler service compiles plugins for runtime proxies.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CompilePlugin | [CompilePluginRequest](#atomix-runtime-v1-CompilePluginRequest) | [CompilePluginResponse](#atomix-runtime-v1-CompilePluginResponse) |  |


<a name="atomix-runtime-v1-PluginRegistry"></a>

### PluginRegistry
The plugin registry service provides plugins for runtime proxies.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| PullPlugin | [PullPluginRequest](#atomix-runtime-v1-PullPluginRequest) | [PullPluginResponse](#atomix-runtime-v1-PullPluginResponse) stream |  |

 



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

