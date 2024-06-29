#!/bin/zsh

#This sets up Apache on your Arch Linux system
#This script is for Arch Linux only

echo "THIS SCRIPT WILL CHANGE YOUR HTTPD CONFIGURATIONS AND VIRTUAL HOSTS. IT WILL ALSO INSTALL APACHE IF IT IS NOT INSTALLED. DO YOU WANT TO CONTINUE? (Y/N)"
read -r response
if [ "$response" != "Y" ]; then
    echo "Exiting..."
    exit 1
fi
if ! command -v httpd &> /dev/null; then
    # Install Apache
    yes | pacman -S httpd -y
fi
if ! command -v migrate &> /dev/null; then
    # Install Golang Migrate from AUR using the yay helper
    yay -S migrate -y
fi
systemctl start httpd
systemctl enable httpd
cp ./.conf/srv.index.html /srv/http/index.html
mkdir /etc/httpd/conf/vhosts
cp ./.conf/xeonlib.org.conf /etc/httpd/conf/vhosts/xeonlib.org
mkdir /srv/xeonlib.org
cp ./.conf/srv.index.html /srv/xeonlib.org/index.html
chown -R root:http /srv/xeonlib.org
rm /etc/httpd/conf/httpd.conf
cp ./.conf/httpd.conf /etc/httpd/conf/httpd.conf
echo "127.0.0.1     xeonlib.org" >> /etc/hosts
systemctl restart httpd
echo "Please provide a connection string to your database. This will be used for migrations:"
read -r conn
echo "
migration_up:
		@read -p \"Enter version: \" v; \\
		migrate -path database/migration/ -database \"$conn\" -verbose up $$v
migration_down:
		@read -p \"Enter version: \" v; \\
		migrate -path database/migration/ -database \"$conn\" -verbose down $$v
migration_fix:
		@read -p \"Enter version: \" v; \\
		migrate -path database/migration/ -database \"$conn\" force $$v
" >> Makefile

