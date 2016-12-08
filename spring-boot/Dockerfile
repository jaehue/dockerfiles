FROM java:openjdk-8u111-jdk

# install mvn
ARG MAVEN_VERSION=3.3.9
RUN mkdir -p /usr/share/maven /usr/share/maven/ref \
  && curl -fsSL http://apache.osuosl.org/maven/maven-3/$MAVEN_VERSION/binaries/apache-maven-$MAVEN_VERSION-bin.tar.gz \
    | tar -xzC /usr/share/maven --strip-components=1 \
  && ln -s /usr/share/maven/bin/mvn /usr/bin/mvn

ADD ./pom.xml /app/pom.xml
WORKDIR /app
RUN mvn clean install -Dmaven.test.skip=true