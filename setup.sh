#/bin/bash
set -e

VERSION="1.9.3"

print_help() {
    echo "Setup script for ContRoll. Installs most recent version of Go, sets environment variables, and compiles everything"
    echo "This script should be run as root\n\n"
    echo "Usage: ./goinstall.sh OPTIONS"
    echo "\nOPTIONS:"
    echo "  --32\t\tInstall 32-bit version"
    echo "  --64\t\tInstall 64-bit version"
}

echo "Setting up ContRoll..."
if [ "$1" = "--32" ]; then
    DFILE="go$VERSION.linux-386.tar.gz"
elif [ "$1" = "--64" ]; then
    DFILE="go$VERSION.linux-amd64.tar.gz"
elif [ "$1" = "--help" ]; then
    print_help
    exit 0
else
    print_help
    exit 1
fi

if [ -d "$HOME/.go" ] || [ -d "$HOME/go" ]; then
    echo "[!]It appears Go is already installed on your system. Exiting..."
    exit 1
fi

echo "[-]Downloading $DFILE ..."
wget https://dl.google.com/go/$DFILE -O /tmp/go.tar.gz || echo "[!]Connection to website failed. Exiting..."

if [ $? -ne 0 ]; then
    echo "[!]Downloading $DFILE failed. Exiting..."
    exit 1
fi

echo "\n[+]Download successfull"
echo "\n[-]Extracting..."
tar -C "$HOME" -xzf /tmp/go.tar.gz
if [ $? -ne 0 ]; then
    echo "[!]Extraction failed. Exiting..."
    exit 1
fi

mv "$HOME/go" "$HOME/.go"
{
    echo '# GoLang'
    echo 'export GOROOT=$HOME/.go'
    echo 'export PATH=$PATH:$GOROOT/bin'
    echo 'export GOPATH=$HOME/go'
    echo 'export PATH=$PATH:$GOPATH/bin'
} >> "$HOME/.bashrc"

mkdir -p $HOME/go/{src,pkg,bin}
echo "\n[+]Go $VERSION installed"
rm -f /tmp/go.tar.gz
$HOME/.go/bin/go build -o ContRoll main.go
##Add ContRoll to /usr/bin??
echo "\n[+]ContRoll successfully installed!"
echo "\n[+]Please run ./ContRoll"
