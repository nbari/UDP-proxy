 [ ![Download](https://api.bintray.com/packages/nbari/UDP-proxy/UDP-proxy/images/download.svg) ](https://bintray.com/nbari/UDP-proxy/UDP-proxy/_latestVersion)

# UDP-proxy

**UDP-proxy**, listens on an UDP port (default :1514) and forwards the traffic
via TCP or UDP to a remote host.

Why ?
=====

Because of the need to have logs per application, stored and rotated on the same
host besides been also been available to send the logs to a remote host, helping
with this the centralization/unification of logs from multiple applications but
still continue debugging locally.


Use case
========

[runit](http://smarden.org/runit/)/[svlogd](http://smarden.org/runit/svlogd.8.html) + [papertrail](https://papertrailapp.com/)

runit is used to keep up and running the application, while keeping logs
in ``log/main/current`` but at the same time is required to send logs to
papertrail.

To do this, ``svlogd`` needs a Config file located in ``log/main/config``
directory, contents of the file could be something like:

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

    UDP-proxy -r logs.papertrailapp.com:61653 -f

This will start **UDP-proxy** listening on port ``127.0.0.1:1514`` and forward via
TCP all request to ``logs.papertrailapp.com:61653``.


DNS proxy
=========

To proxy all request to dns.watch:

    UDP-proxy -b 127.0.0.1:5253 -r resolver1.dns.watch:53

Test the proxy using dig:

    dig @127.0.0.1 -p 5253 github.com mx

Iperf
=====

Listen on localhost port 5001:

    UDP-proxy -r localhost:5001 -d

Run iperf server:

    iperf -s -u

Connect the client:

    iperf -u -c 127.0.0.1 -p 1514 -d
