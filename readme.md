# WoodKV
WoodKV is a simple in-memory database, based on skiplist. The data model is similar to map structure in C++.   
It provides Get, Put, Delete, PrefixScan interfaces.
## Quick Start
```
$ cd main
$ go run server.go
```
In another window, run client.
```
$ go run client.go
```
**Hint: By default, the key can not exceed 8 bytes and the value can not exceed 256 bytes.**
## Examples
```
>put username Zhangsan
succeed.
>put password 123456
succeed.
>get username
(Zhangsan



                 )
>prefixscan 2
password username
>delete username
succeed.
>get username
[error] 101: 找不到
```

## TODO 

- [ ] testing
- [ ] persistence 

## Why SkipList?
Simple