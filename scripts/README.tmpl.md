# until it works

Retry some command until it works.

```
$ untilitworks curl http://mywebsite.com
```

<img width="376" src="https://cloud.githubusercontent.com/assets/1189716/12663176/b006cace-c5f3-11e5-9654-271063ac9510.png">

## usage

Exponential backoff (with random jitter):

```
$ untilitworks -retry exponentially exit 1
untilitworks: it failed! retrying in 583.112038ms
untilitworks: it failed! retrying in 1.447272331s
untilitworks: it failed! retrying in 1.819838229s
untilitworks: it failed! retrying in 2.514838034s
untilitworks: it failed! retrying in 9.191532459s
```

Constant backoff:
```
$ untilitworks -retry constant exit 1
untilitworks: it failed! retrying in 1s
untilitworks: it failed! retrying in 1s
untilitworks: it failed! retrying in 1s
```


## installation

### linux

```bash
wget -qO- https://github.com/aybabtme/untilitworks/releases/download/{{.version}}/untilitworks_linux.tar.gz | tar xvz
```

### darwin

```bash
wget -qO- https://github.com/aybabtme/untilitworks/releases/download/{{.version}}/untilitworks_darwin.tar.gz | tar xvz
```


## halp!

```bash
Usage of untilitworks:
  -exp.cap duration
    	max time to backoff for (default 30s)
  -exp.factor float
    	backoff factor for exponential retries (default 2)
  -q	whether to suppress the command's output
  -retry string
    	retry type, one of 'c'/'constantly' or 'e'/'exponentially' (default "constantly")
  -sleep duration
    	how long to sleep between retries (base duration for exponential) (default 1s)
```


## license

MIT
