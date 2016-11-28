FROM jaehue/hdfs-base
MAINTAINER Jaehue Jang <jaehue@jang.io>

# Setup env
ENV HIVE_VERSION 2.1.0

USER root

# download hive
RUN wget -q -O - http://mirror.navercorp.com/apache/hive/hive-$HIVE_VERSION/apache-hive-$HIVE_VERSION-bin.tar.gz | tar -xzf - -C /usr/local \
 && ln -s /usr/local/apache-hive-$HIVE_VERSION-bin /usr/local/apache-hive-bin

# Setup permissions and ownership
RUN chown -R hadoop:hadoop /usr/local/apache-hive-$HIVE_VERSION-bin \
 && chmod -R 775 /usr/local/apache-hive-$HIVE_VERSION-bin/bin /usr/local/apache-hive-$HIVE_VERSION-bin/scripts

USER hadoop

# download mysql jdbc connector
RUN wget -q -O - http://dev.mysql.com/get/Downloads/Connector-J/mysql-connector-java-5.1.40.tar.gz | tar -xzf - -C /usr/local/apache-hive-bin/lib

EXPOSE 10000
EXPOSE 10002