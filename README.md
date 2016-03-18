# UDP-proxy

**UDP-proxy**, listen on an UDP port (default 1514) and forwards the traffic to a
TCP or UDP host.

Why ?
=====

Because of the need to have logs per application, stored and rotated on the same
host besides been also been available to send the logs to a remote host, helping
with this the centralization/unification of logs from multiple applications but
still continue debugging locally.


Use case
========

[runit](http://smarden.org/runit/)/[svlogd](http://smarden.org/runit/svlogd.8.html) + [papertrail](https://papertrailapp.com/)

runit is using to keep up and running the application, while keeping logs in ``log/main/current`` but at the same time is required to send logs to papertrail.

To do this, ``svlogd`` needs a Config file located in ``log/main/config`` directory, contents of the file could be something like:

    s1000000
    n10
    N5
    t86400
    u127.0.0.1:1514

Notice the last line ``u127.0.0.1:1514``, from the svlogd.8 man page:

```text
ua.b.c.d[:port]
tells svlogd to transmit the first len characters of selected log messages to
the IP address a.b.c.d, port number port. If port isn’t set, the default port
for syslog is used (514). len can be set through the -l option, see below. If
svlogd has trouble sending udp packets, it writes error messages to the log
directory. Attention: logging through udp is unreliable, and should be used in
private networks only.
```

Setting **UDP-proxy** to forward logs to **papertrail** is straight forward:

    UDP-proxy -f logs.papertrailapp.com:61653

This will start **UDP-proxy** listening on port ``127.0.0.1:1514`` and forward via
TCP all request to ``logs.papertrailapp.com:61653``.



Testing
=======

You can use the included test client ``client/main`` and
[netcat](https://en.wikipedia.org/wiki/Netcat) to receive the
requests, first start **UDP-proxy**:

    UDP-proxy -f localhost:9090 -d

Next netcat to keep listening on port 9090:

    nc -lk 9090

And at the end start the client:

    $ cd client
    $ go run main.go
