# HTTP echo Docker image

    $ docker run --net=host -it hekonsek/http-echo
    Started echo server on http://localhost:8080 .

    $curl http://localhost:8080/you
    Hi there, I love you!