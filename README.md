<<<<<<< HEAD
# ![pos](http://s3.amazonaws.com/dgvstuff/pos.png) omg8583 [![Build Status](https://drone.io/github.com/danielvargas/omg8583/status.png)](https://drone.io/github.com/danielvargas/omg8583/latest)
**[ISO-8583](http://en.wikipedia.org/wiki/ISO_8583) Go lightweight library**

=======
# omg8583 [![Build Status](https://drone.io/github.com/danielvargas/omg8583/status.png)](https://drone.io/github.com/danielvargas/omg8583/latest)
>>>>>>> 20eae19a6038e7e56ee7596ef3e77cdf5116fddf

    import "github.com/danielvargas/omg8583"
    

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
