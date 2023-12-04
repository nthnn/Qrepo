git clone https://github.com/nthnn/Qrepo.git
cd Qrepo
chmod -R 777 build.sh
./build.sh
mv dist/qrepo /usr/local/bin/qrepo
cd .. && sudo rm -rf Qrepo