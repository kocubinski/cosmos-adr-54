# Compatibility Sandbox

This is a sandbox for showing trivialized examples of compatibility issues between protobuf versions and SDK modules 
as they relate to the problems and solutions described in ADR-54.

## Key Concepts
- [Unknown Fields](https://developers.google.com/protocol-buffers/docs/proto3#unknowns) 
- [ADR-20 on Unknown Fields](https://github.com/cosmos/cosmos-sdk/blob/aaronc/adr-proto-go-module/docs/architecture/adr-020-protobuf-transaction-encoding.md#unknown-field-filtering )
- [Any Types](https://developers.google.com/protocol-buffers/docs/proto3#any)

### Forward / Backward compatibility

Forward compatibility is a design characteristic that allows a system to accept input intended for a later
version of itself. This is the opposite of backwards compatibility, which is the ability to accept input from
a previous version of the system.

A system supports forward compatibility if a module that complies with earlier versions can gracefully process
input designed for later versions of the system, ignoring new parts which it does not understand.  A system
supports backward compatibility if a module that complies with later versions can gracefully process input
designed for earlier versions, sensibly driving newer features from legacy input. In our case, input means raw
bytes (e.g. a transaction) and module means an SDK module which is also a standalone golang module.

For the Cosmos SDK the interplay between a single state-machine module (at a given version) and the API
specification (at a possibly different version) contextualizes forward/backward compatibility. The API
surface, specified by protobuf (and realized by generated golang code) defines the input format by which raw
bytes are translated into concrete domain specific objects that a state-machine module implementing the
Cosmos SDK's API operates on.  These domain specific objects are, in our case, **part of the API specification**.

Therefore, when talking about forward/backward compatibility we are tangibly talking about a **compatibility
matrix between varying versions of the API specification and SDK state-machine modules.**
