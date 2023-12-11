# an apache Apache tomcat 7.0.0 CVE-2017-12615
FROM java:6b32

RUN wget http://archive.apache.org/dist/tomcat/tomcat-7/v7.0.11/bin/apache-tomcat-7.0.11.tar.gz && \
    tar -xvf apache-tomcat-7.0.11.tar.gz && \
    mv apache-tomcat-7.0.11 /usr/local/tomcat && \
    rm -rf apache-tomcat-7.0.11.tar.gz

# prepare tomcat vulnerable version
RUN mkdir -p /usr/local/tomcat/webapps/ROOT
RUN echo "Hello, World" > /usr/local/tomcat/webapps/ROOT/index.html

# exploit

# add in conf/web.xml param-name "readonly" to false, use "<load-on-startup>3</load-on-startup>" 
COPY apache-tomcat-web.xml /usr/local/tomcat/conf/web.xml
RUN cat /usr/local/tomcat/conf/web.xml
# RUN cat /usr/local/tomcat/webapps/ROOT/WEB-INF/web.xml

EXPOSE 8080

CMD ["/usr/local/tomcat/bin/catalina.sh", "run"]



