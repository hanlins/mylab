server {
	listen 80; 		# ipv4
	listen [::]:80; 	# ipv6

	server_name mylab.com; 	# vhost

	location / {
		#root /root/mylab/;
		proxy_pass http://127.0.0.1:8001;
	}

	#location /static/ {
		
}

