This is a repo for testing a website hosted at 'mylab.com'

Use
`go build`
to build the go server (which listens to port 8001).

Then nginx will do reverse proxy to redirect request from port 80 to 8001.
Notice nginx config file should be placed in
/etc/nginx/sites-available/mylab.com, then create a soft link in the other
directory
`ln -s /etc/nginx/sites-enabled/mylab.com /etc/nginx/sites-available/mylab.com`

