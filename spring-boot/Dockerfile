FROM jaehue/maven

ADD ./pom.xml /app/pom.xml
WORKDIR /app
RUN mvn clean install -Dmaven.test.skip=true