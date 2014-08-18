#scratch2hue

_library for manipulating the hue through the scratch._

## Install
### Go
Refer to [golang.org](http://golang.org/doc/install)

### gom
Refer to [gom](https://github.com/mattn/gom)

### scratch2hue
```bash
$ go get github.com/135yshr/scratch2hue
$ gom install
$ cd sample
$ go build -o scratch2hue
```

## How to use
1. start the scratch
1. open the hue.sb
1. [Enable Remote Sensor Connections](http://wiki.scratch.mit.edu/wiki/Remote_Sensor_Connections)
1. start the scratch2hue

### command to start the scratch2hue
```bash
$ ./scratch2hue -ip <bridge ip address>
```

## License
MIT:https://github.com/135yshr/scratch2hue/blob/master/LICENSE

## Revision History
### beta_1.1
* changed of function name from scratch2hue.Anction to scratch2hue.Action.
* changed of import from "scratch2hue" to "github.com/135yshr/scratch2hue"