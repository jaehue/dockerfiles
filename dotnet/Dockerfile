FROM microsoft/dotnet:onbuild

COPY ./project.json /app/
WORKDIR /app
RUN ["dotnet", "restore"]