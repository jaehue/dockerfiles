FROM java:openjdk-8u111-jdk
MAINTAINER Jaehue Jang <jaehue@jang.io>

# Setup env
ENV HADOOP_USER hdfs
ENV JAVA_HOME /usr/lib/jvm/java-8-openjdk-amd64
ENV HADOOP_VERSION 2.7.3

USER root

# download hadoop
RUN wget -q -O - http://apache.mirror.cdnetworks.com/hadoop/common/hadoop-$HADOOP_VERSION/hadoop-$HADOOP_VERSION.tar.gz | tar -xzf - -C /usr/local \
 && ln -s /usr/local/hadoop-$HADOOP_VERSION /usr/local/hadoop \
 && groupadd -r hadoop \
 && groupadd -r $HADOOP_USER && useradd -r -g $HADOOP_USER -G hadoop $HADOOP_USER

USER $HADOOP_USER