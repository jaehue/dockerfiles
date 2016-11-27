FROM jaehue/hdfs-base
MAINTAINER Jaehue Jang <jaehue@jang.io>

# Setup env
ENV JAVA_HOME /usr/lib/jvm/java-8-openjdk-amd64
ENV HIVE_VERSION 2.1.0
ENV HADOOP_HOME /usr/local/hadoop
ENV HADOOP_PREFIX /usr/local/hadoop
ENV HIVE_USER hive

USER root

RUN wget -q -O - http://mirror.navercorp.com/apache/hive/hive-$HIVE_VERSION/apache-hive-$HIVE_VERSION-bin.tar.gz | tar -xzf - -C /usr/local \
 && ln -s /usr/local/apache-hive-$HIVE_VERSION-bin /usr/local/apache-hive-bin \
 && chown -R $HIVE_USER:hive /usr/local/apache-hive-$HIVE_VERSION-bin \
 && chmod -R 775 /usr/local/apache-hive-$HIVE_VERSION-bin/bin /usr/local/apache-hive-$HIVE_VERSION-bin/scripts

USER $HIVE_USER