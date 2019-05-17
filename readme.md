# HTTP echo Docker image

## Usage

Start Docker image:

    $ docker run --net=host -it hekonsek/http-echo
    Started echo server on http://localhost:8080 .

Use HTTP client to receive an answer:

    $ curl http://localhost:8080/you
    Hi there, I love you!