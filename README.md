moremon
=======

tiny little linux monitoring in golang and websockets with JavaScript example
based on jQuery/Flot. Features:

* monitors load

Live demo
---------

Live demo is running on Red Hat OpenShift hosted platform cloud. Give the
example some time to load, the application is usually stopped when you click
the link: http://moremon-lzap.rhcloud.com

Building
--------

    git clone https://github.com/lzap/moremon.git
    cd moremon
    go get github.com/gorilla/websocket
    go get github.com/c9s/goprocinfo/linux
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
