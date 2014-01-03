# omg8583

    import "github.com/danielvargas/omg8583"
    
[ISO-8583](http://en.wikipedia.org/wiki/ISO_8583) Go lightweight library.

![http://s3.amazonaws.com/dgvstuff/vx.gif](http://s3.amazonaws.com/dgvstuff/vx.gif)


### Index

* [func  Pack](#func--pack)
* [func  Unpack](#func--unpack)


#### func  [Pack](#pack)

```go
func Pack(isomsg map[int]string) (string, error)
```

#### func  [Unpack](#unpack)

```go
func Unpack(msg string) (isomsg map[int]string, err error)
```
--
**omg8583** is licensed under BSD license, Copyright (c) 2013 Daniel Vargas.