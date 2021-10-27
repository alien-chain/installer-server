# installer-server

## Create a systemd process for alieninstall
First, create some necessary files

```
sudo mkdir -p /var/log/alieninstall 
sudo touch /var/log/alieninstall/alieninstall.log 
sudo touch /var/log/alieninstall/alieninstall_error.log 
sudo touch /etc/systemd/system/alieninstall.service
```

Second, edit the systemd service file for alieninstall.

```
sudo nano /etc/systemd/system/alieninstall.service
```

Next, add the following configuration to it:

```
Description=alieninstall daemon
After=network-online.target
[Service]
User=first_user
ExecStart=/home/first_user/go/bin/alieninstall start --home=/home/first_user/.alieninstall
WorkingDirectory=/home/first_user/go/bin
StandardOutput=file:/var/log/alieninstall/alieninstall.log
StandardError=file:/var/log/alieninstall/alieninstall_error.log
Restart=always
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target
```

Then, enabled it to run all the time even after it reboots.

```
sudo systemctl enable alieninstall.service
```

Start the process.

```
sudo systemctl start alieninstall.service
```

Lastly, I can see the logs of the new service like this:
```
sudo journalctl -u alieninstall -f
```

Nginx Reverse Proxy

Install Nginx

```
apt-get install nginx
```

Add entry & enable. Create this file in sites-available, create link to sites-enabled. Delete default site link.

```
server {
        listen  80;
        server_name some.domain.org;
        location / {
            proxy_set_header    Host $host;
            proxy_set_header    X-Real-IP   $remote_addr;
            proxy_set_header    X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_pass  http://127.0.0.1:8081;
        }
}
```


