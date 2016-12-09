FROM jaehue/spring-boot

WORKDIR /app
ADD ./pom.xml /app/pom.xml
RUN mvn clean install -Dmaven.test.skip=true

ADD ./package.json /app/package.json
RUN npm install