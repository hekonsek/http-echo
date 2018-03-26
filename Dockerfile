FROM fedora:27

ADD http-echo /home/fedora/http-echo

EXPOSE 8080

ENTRYPOINT ["/home/fedora/http-echo"]