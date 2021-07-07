# Install dependencies
FROM debian:latest
RUN apt-get update 
RUN apt-get install -y curl git wget unzip libgconf-2-4 gdb libstdc++6 libglu1-mesa fonts-droid-fallback lib32stdc++6 psmisc
RUN apt-get clean

# Clone the flutter repo
RUN git clone https://github.com/flutter/flutter.git /usr/local/flutter

# Set flutter path
ENV PATH="/usr/local/flutter/bin:/usr/local/flutter/bin/cache/dart-sdk/bin:${PATH}"

# Enable flutter web
RUN flutter channel stable
RUN flutter upgrade
RUN flutter config --enable-web
RUN flutter config --enable-windows-desktop
RUN flutter config --enable-macos-desktop
RUN flutter config --enable-linux-desktop

# Run flutter doctor
RUN flutter doctor -v

# Setup Web SDK
RUN mkdir /app
WORKDIR /app
RUN flutter create myapp
WORKDIR /app/myapp
RUN flutter build web
WORKDIR /
RUN rm -rf /app
