# empty-room
Service that queries empty rooms @ NJIT

## Site
- Site: `http://empty-room.xyz` 

## :microscope: Technologies
- Languages: `go`, `css`, `html`, `python3`
- Packages: `nginx`,`python-pandas`
- Technologies: `make`

## Infra
- [Golang + Nginx Setup](https://github.com/hahdookin/cs490/blob/main/util/dep.sh) 
- [Infrastructure Confs](https://github.com/hahdookin/cs490/tree/main/infra)
- [Nginx Specific Confs](https://gitlab.com/del-repos/spaceZoo/-/blob/main/infra/systemctl-service.conf)

## :microscope: Directory Explanation

| Directory                             | Explanation
| :---:                                 | :---
| [conf](confs)                         | Server Configurations
| [data-processing](data-processing)    | Processes data inside the data-processing/data-collection directory
| [src/](src)                           | Source Code for the web app to get everything set

## :books: Resources
- [Install Ansible OSERROR Python Fix](https://stackoverflow.com/questions/54778630/could-not-install-packages-due-to-an-environmenterror-errno-2-no-such-file-or)
- [Course Catalog](https://myhub.njit.edu/BannerExtensibility/customPage/page/stuRegCrseSched)
- [Golang: Build from source](https://go.dev/doc/install)
- [MySQL New User](https://askubuntu.com/questions/1322175/not-allowed-to-create-user-with-grant)
- [Avoid SQL Injection](https://go.dev/doc/database/sql-injection)
- [nginx proxy_pass server set_up](https://serverfault.com/questions/598202/make-nginx-to-pass-hostname-of-the-upstream-when-reverseproxying)
- [OG Go Digitalocean Service](https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-using-nginx-on-ubuntu-18-04#conclusion)
- [Lets encrypt nginx: ubuntu 20.04](https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-ubuntu-20-04)