# an apache Apache version 2.2.15 container
FROM ubuntu:14.04

RUN apt-get update && apt-get install wget build-essential -y

RUN wget https://archive.apache.org/dist/httpd/httpd-2.2.15.tar.gz && \
    tar -xvf httpd-2.2.15.tar.gz && \
    cd httpd-2.2.15 && \
    ./configure --prefix=/usr/local/apache2 --enable-so --build=x86_64-linux-gnu && \
    make && \
    make install && \
    rm -rf /httpd-2.2.15 && \
    rm -rf /httpd-2.2.15.tar.gz

EXPOSE 80

CMD ["/usr/local/apache2/bin/httpd", "-D", "FOREGROUND"]
