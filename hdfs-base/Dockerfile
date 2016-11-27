FROM java:openjdk-8u111-jdk
MAINTAINER Jaehue Jang <jaehue@jang.io>

# Setup env
ENV HADOOP_USER hdfs
ENV HADOOP_VERSION 2.7.3

USER root
RUN groupadd -r hadoop \
 && groupadd -r $HADOOP_USER && useradd -r -g $HADOOP_USER -G hadoop $HADOOP_USER

# download hadoop
RUN wget -q -O - http://apache.mirror.cdnetworks.com/hadoop/common/hadoop-$HADOOP_VERSION/hadoop-$HADOOP_VERSION.tar.gz | tar -xzf - -C /usr/local \
 && ln -s /usr/local/hadoop-$HADOOP_VERSION /usr/local/hadoop

# Setup permissions and ownership
RUN chown -R $HADOOP_USER:hadoop /usr/local/hadoop-$HADOOP_VERSION \
 && chmod -R 755 /usr/local/hadoop-$HADOOP_VERSION/etc/hadoop \
 && chmod -R 755 /usr/local/hadoop-$HADOOP_VERSION/bin

USER $HADOOP_USER