FROM java:openjdk-8u111-jdk
MAINTAINER Jaehue Jang <jaehue@jang.io>

# Setup env
ENV JAVA_HOME /usr/lib/jvm/java-8-openjdk-amd64
ENV HADOOP_VERSION 2.7.3

# download hadoop
RUN wget -q -O - http://apache.mirror.cdnetworks.com/hadoop/common/hadoop-$HADOOP_VERSION/hadoop-$HADOOP_VERSION.tar.gz | tar -xzf - -C /usr/local \
 && ln -s /usr/local/hadoop-$HADOOP_VERSION /usr/local/hadoop