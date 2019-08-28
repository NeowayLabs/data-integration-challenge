FROM openjdk:8-jdk-alpine

RUN mkdir -p /pdv-service/update-dir

COPY src /pdv-service/src
COPY mvnw /pdv-service/mvnw
COPY .mvn /pdv-service/.mvn
COPY pom.xml /pdv-service/pom.xml

RUN cd /pdv-service && sh mvnw clean install

ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/pdv-service/target/pdv-service.jar"]
