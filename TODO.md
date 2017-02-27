# Nerdz Core TODO list

Nerdz Core is designed to offer a single entrypoint for each service offered by Nerdz, including the API and Web.

## What has been done in Nerdz API

+ Created types (ORM model)
+ Fetch comments and posts (with related options: from friends only, in a language only and these options can be mixed).
+ Add/Delete/Edit comment/post
+ Rereiving user information (numeric (fast) or complete)
+ Implement the messages interfaces for pm -> write tests
+ Add a method for every user action (follow, update post/comment, create things and so on)
+ Tests for every method

## What has been done in Nerdz Core
+ Branched Core from Nerdz API

## What needs to be done
+ Finish moving Db types/logic into the `db` package (tests)
+ Provide a full Protocol Buffers description of the protocol Nerdz Core wishes to expose to the world
+ Define an RPC protocol using gRPC
+ Provide TLS-based authentication
+ Provide 100% testing coverage 

A big thanks to Paolo for its long-standing commitment to the development of Nerdz API.

## What has been contributed to the [gorm](https://github.com/jinzhu/gorm/) project:
- [Add support for primary key different from id](https://github.com/jinzhu/gorm/pull/85)
- [Add support to fields with double quotes](https://github.com/jinzhu/gorm/pull/105)
- [Add default values support](https://github.com/jinzhu/gorm/pull/279)
