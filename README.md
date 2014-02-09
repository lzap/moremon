moremon
=======

tiny little linux monitoring in golang and websockets with JavaScript example
based on jQuery/Flot. Features:

* monitors load

Building
--------

    git clone https://github.com/lzap/moremon.git
    cd moremon
    go get github.com/gorilla/websocket
    go get github.com/alecthomas/gozmq
    go build
    ./moremon

Using
-----

Just visit localhost:8080

    ./moremon -h
    Usage of ./moremon:
      -addr=":8080": http service address (default: localhost:8080)
      -history=300: history size in samples (default: 300)
      -update=1: update interval in seconds (default: 1)

Todo
----

* Integration with Foreman as a plugin.
