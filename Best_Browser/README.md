# Best Browser

## Installation guide:
* [Download](https://golang.org/dl/) GO 
* [Download](https://www.nginx.com/resources/admin-guide/installing-nginx-open-source/) Nginx
* Copy the following nginx-server.conf
```
http {
    upstream bestbrowser {
        server localhost:1234;
    }
    
    server {
        listen 80;

        location / {
            proxy_pass http://bestbrowser;
        }
    }
}
events {}
```
* [Download](https://www.postgresql.org/download/) PostgresSQL
* Create a user "bb" with password "bb"
* Create database:
```
CREATE DATABASE best_browser;
```
* Create table:
```
CREATE TABLE VOTES ( FIREFOX INT ,CHROM INT, EXPLORER INT);
INSERT INTO VOTES (firefox,chrom,explorer) VALUES (0,0,0);
```
* Run the command 
```
go run /path/to/files/main.go
```
* In your browser navigate to "localhost"
