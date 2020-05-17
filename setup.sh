sudo yum install git
git clone https://github.com/kahgeh/osserv.git

curl -fSL https://dl.google.com/go/go1.14.3.linux-amd64.tar.gz --output go.tar.gz
sudo tar -C /usr/local -xzf go.tar.gz
export PATH=$PATH:/usr/local/go/bin