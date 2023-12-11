#!/bin/bash

# tomcat
# cd dockerfiles/tomcat
# docker build -t vuln-apache-tomcat-7.0.0 -f apache-tomcat-7.0.0.dockerfile --progress plain .
TOMCAT=$(docker run -d -p8080:8080 vuln-apache-tomcat-7.0.0)

# php

# cd $(mktemp -d)
# git clone https://github.com/neex/phuip-fpizdam
# cd phuip-fpizdam/reproducer/
# docker build -t reproduce-cve-2019-11043 .
PHP=$(docker run -d -p 8081:80 --name reproduce-cve-2019-11043 reproduce-cve-2019-11043)

# cd $(mktemp -d)
# git clone https://github.com/kozmer/log4j-shell-poc.git
# cd log4j-shell-poc
# # replace in Dockerfile "tomcat:8.0.36-jre8" with "FROM tomcat:8.0.35-jre8"
# sed -i '' 's/tomcat:8.0.36-jre8/tomcat:8.0.35-jre8/g;' Dockerfile
# docker build -t log4j-shell-poc .
# LOG4J=$(docker run -d -p 8082:80 log4j-shell-poc)


trap "docker stop $TOMCAT $PHP $LOG4J && docker rm $TOMCAT $PHP $LOG4J" EXIT

echo "###### READY ######"

read -r -d '' _ </dev/tty