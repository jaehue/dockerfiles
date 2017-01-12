FROM microsoft/dotnet:1.1.0-sdk-projectjson

COPY ./project.json /app/
WORKDIR /app
RUN ["dotnet", "restore"]