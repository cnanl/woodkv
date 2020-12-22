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
>put userxyz test111111111111111111
succeed.
>put user11 zzzx
succeed.
>put userzxc mxxxxxx
succeed.
>prefixscan user 3 
user11   username userxyz
```

## TODO 

- [ ] testing
- [ ] persistence 

## Why SkipList?
Simple