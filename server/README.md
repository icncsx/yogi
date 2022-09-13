

if the requester (client) has its IDL field set to optional,
then ther


required
- required field does not have the isset attribute


<!-- - if field is default/required and the received field is empty (not set) the remote end (receiver) gets an empty struct (default value) -->

default
- has isset attribute


<!-- - if not assigned, its value is the default value
- if field is default/required and the received field is empty (not set) the remote end (receiver) gets an empty struct (default value) -->

optional
- has isset attribute


<!-- - if isset is false, value will be lost after transmission (or seriaization/deserialization)
- when the field is optional, when the struct field written is nil, the far end must be nil -->