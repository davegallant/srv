#!/usr/bin/env bash

VERSION='0.1.1'; \
sudo curl --progress-bar \
-L "https://github.com/davegallant/srv/releases/download/v${VERSION}/srv_${VERSION}_$(uname -s)_x86_64.tar.gz" | \
sudo tar -C /usr/bin --overwrite -xvzf - srv
