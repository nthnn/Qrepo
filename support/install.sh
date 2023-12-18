if [ "$(uname)" == "Darwin" ]; then
    (brew list golang || brew install golang) > /dev/null 2>&1
    (brew list git || brew install git) > /dev/null 2>&1
elif [ "$(uname)" == "Linux" ]; then
    sudo apt-get -y install golang > /dev/null 2>&1
    sudo apt-get -y install git > /dev/null 2>&1
fi

git clone --quiet --depth 1 https://github.com/nthnn/Qrepo.git && cd Qrepo
chmod -R 777 build.sh && ./build.sh
rm -i /usr/local/bin/qrepo && mv dist/qrepo /usr/local/bin/qrepo && rm -rf ../Qrepo