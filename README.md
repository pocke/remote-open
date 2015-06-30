remote-open
=============

open/start/xdg-open over TCP.


Installation
-------------

```sh
go get github.com/pocke/remote-open/...
```

Or download from https://github.com/pocke/remote-open/releases/latest.

Usage
------

`port` default is 2489.

### Server

```sh
remote-opend --port 1234
```

### Client


```sh
remote-open --port 1234 --host '192.168.x.x' 'http://example.com'
```

`http://example.com` is opened by browser on Server.



Configuration
----------------

### Server

`~/.config/remote-opend.toml`

```toml
port = 1234
```

### Client

`~/.config/remote-open.toml`

```toml
port = 1234
host = '192.168.x.x'
```
