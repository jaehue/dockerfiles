FROM microsoft/dotnet:1.0.1-sdk-projectjson

COPY ./project.json /app/
WORKDIR /app
RUN ["dotnet", "restore"]
