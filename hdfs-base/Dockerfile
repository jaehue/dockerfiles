FROM java:openjdk-8u111-jdk
MAINTAINER Jaehue Jang <jaehue@jang.io>

# Setup env
ENV HADOOP_VERSION 2.7.3

USER root
RUN groupadd -r hadoop && useradd -r -g hadoop -G hadoop hadoop

# download hadoop
RUN wget -q -O - http://apache.mirror.cdnetworks.com/hadoop/common/hadoop-$HADOOP_VERSION/hadoop-$HADOOP_VERSION.tar.gz | tar -xzf - -C /usr/local \
 && ln -s /usr/local/hadoop-$HADOOP_VERSION /usr/local/hadoop

# Setup permissions and ownership
RUN chown -R hadoop:hadoop /usr/local/hadoop-$HADOOP_VERSION \
 && chmod -R 755 /usr/local/hadoop-$HADOOP_VERSION/etc/hadoop \
 && chmod -R 755 /usr/local/hadoop-$HADOOP_VERSION/bin

USER hadoop

# Namenode
EXPOSE 8020 8022 50070 50470 50090 50495
# DataNode
EXPOSE 1006 1004 50010 50020 50075 50475
# Journal Node WebHDFS
EXPOSE 8485 8480 14000