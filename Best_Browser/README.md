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
* Create a user "bb" and password "bb"
* Create a database named "best_browser"
```
CREATE DATABASE best_browser;
```
* Create a table named "votes":
```
CREATE TABLE VOTES ( FIREFOX INT ,CHROM INT, EXPLORER INT);
INSERT INTO VOTES (firefox,chrom,explorer) VALUES (0,0,0);
```
* Run the command 
```
go run /path/to/files/main.go
```
* In your browser navigate to "localost"
