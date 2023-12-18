echo "[\033[92m+\033[0m] Installing Qrepo..."

echo "[\033[92m+\033[0m] Checking up Go compiler and git..."
if [ "$(uname)" == "Darwin" ]; then
    (brew list golang || brew install golang) > /dev/null 2>&1
    (brew list git || brew install git) > /dev/null 2>&1
elif [ "$(uname)" == "Linux" ]; then
    sudo apt-get -y install golang > /dev/null 2>&1
    sudo apt-get -y install git > /dev/null 2>&1
fi

echo "[\033[92m+\033[0m] Cloning source files from GitHub..."
git clone --quiet --depth 1 https://github.com/nthnn/Qrepo.git && cd Qrepo
chmod -R 777 build.sh && ./build.sh

if[ -e "/usr/local/bin/qrepo" ]; then
    echo "[\033[91m-\033[0m] Removing previously installed Qrepo..."
    rm -rf /usr/local/bin/qrepo
fi

echo "[\033[92m+\033[0m] Installing Qrepo to default bin path..."
mv dist/qrepo /usr/local/bin/qrepo && rm -rf ../Qrepo

echo "[\033[92m+\033[0m] Successfully \033[36installed\033[0m!"