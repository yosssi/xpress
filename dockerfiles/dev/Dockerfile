# yosssi/xpress-dev
#
# VERSION 1.0.0

# use the ubuntu base image provided by dotCloud
FROM ubuntu

MAINTAINER Keiji Yoshida, yoshida.keiji.84@gmail.com

# make sure the package repository is up to date
RUN echo "deb http://archive.ubuntu.com/ubuntu precise main universe" > /etc/apt/sources.list
RUN apt-get update

# install libraries
RUN apt-get install -y software-properties-common python-software-properties

# install make curl
RUN apt-get install -y make curl

# install git
RUN add-apt-repository -y ppa:git-core/ppa
RUN apt-get update
RUN apt-get install -y git
RUN git config --global user.email "yoshida.keiji.84@gmail.com"
RUN git config --global user.name "yosssi"
RUN git config --global core.editor "/usr/bin/vi"
RUN git config --global color.ui "true"

# install Go 1.2.1
RUN curl -o /usr/local/go1.2.1.linux-amd64.tar.gz https://go.googlecode.com/files/go1.2.1.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf /usr/local/go1.2.1.linux-amd64.tar.gz
RUN rm /usr/local/go1.2.1.linux-amd64.tar.gz
ENV GOROOT /usr/local/go
ENV GOPATH /host/go
ENV PATH $PATH:$GOPATH/bin:$GOROOT:$GOROOT/bin

# install Java 7
RUN curl -o /usr/local/lib/jre-7u51-linux-x64.tar.gz https://s3-ap-northeast-1.amazonaws.com/yosssi/java/jre-7u51-linux-x64.gz
RUN tar -C /usr/local/lib -xzf /usr/local/lib/jre-7u51-linux-x64.tar.gz
RUN rm /usr/local/lib/jre-7u51-linux-x64.tar.gz
ENV JAVA_HOME /usr/local/lib/jre1.7.0_51
ENV PATH $PATH:$JAVA_HOME/bin

# install Elasticsearch 1.0.1
RUN curl -o /usr/local/lib/elasticsearch-1.1.0.tar.gz https://download.elasticsearch.org/elasticsearch/elasticsearch/elasticsearch-1.1.0.tar.gz
RUN tar -C /usr/local/lib -xzf /usr/local/lib/elasticsearch-1.1.0.tar.gz
RUN rm /usr/local/lib/elasticsearch-1.1.0.tar.gz
ENV PATH $PATH:/usr/local/lib/elasticsearch-1.1.0/bin

# install Node.js 0.10.26
RUN curl -o /usr/local/lib/node-v0.10.26-linux-x64.tar.gz http://nodejs.org/dist/v0.10.26/node-v0.10.26-linux-x64.tar.gz
RUN tar -C /usr/local/lib -xzf /usr/local/lib/node-v0.10.26-linux-x64.tar.gz
RUN rm /usr/local/lib/node-v0.10.26-linux-x64.tar.gz
ENV NODE_HOME /usr/local/lib/node-v0.10.26-linux-x64
ENV PATH $PATH:$NODE_HOME/bin

# install Bower
RUN npm install -g bower

# install Stylus
RUN npm install -g stylus

# install UglifyJS2
RUN npm install -g uglify-js

# install UglifyCSS
RUN npm install -g uglifycss

# install Goat
RUN curl -o /usr/local/bin/goat https://s3-ap-northeast-1.amazonaws.com/yosssi/goat/linux_amd64/goat
RUN chmod +x /usr/local/bin/goat

# install gcc.
RUN apt-get install -y gcc

# install Redis 2.8.8
RUN curl -o /usr/local/lib/redis-2.8.8.tar.gz http://download.redis.io/releases/redis-2.8.8.tar.gz
RUN tar -C /usr/local/lib -xzf /usr/local/lib/redis-2.8.8.tar.gz
RUN rm /usr/local/lib/redis-2.8.8.tar.gz
RUN cd /usr/local/lib/redis-2.8.8; make
ENV REDIS_HOME /usr/local/lib/redis-2.8.8
ENV PATH $PATH:$REDIS_HOME/src
