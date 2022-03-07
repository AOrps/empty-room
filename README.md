# empty-room
Service that queries empty rooms @ NJIT

## Site
- Site: `http://empty-room.xyz` 
- API Endpoint: `https://ec2-18-119-118-48.us-east-2.compute.amazonaws.com`


## :tram: Infrastructure
```txt
.------------.             .-------------.
| Front-end  |   <----->   | API Service |
'------------'             '-------------'
```

## :microscope: Technologies
- Languages: `go`, `css`, `html`
- Packages: `nginx`,

## Infra
- [Golang + Nginx Setup](https://github.com/hahdookin/cs490/blob/main/util/dep.sh) 
- [Infrastructure Confs](https://github.com/hahdookin/cs490/tree/main/infra)
- [Nginx Specific Confs](https://gitlab.com/del-repos/spaceZoo/-/blob/main/infra/systemctl-service.conf)

## :microscope: Directory Explanation

| Directory                             | Explanation
| :---:                                 | :---
| [src/front-end](src/front-end)        | Golang/Bootstrap Front-end
| [src/api](src/api)                    | Endpoint to get information about empty-rooms @ NJIT
| [data-processing](data-processing)    | Data Processing from data-collection dir

## :books: Resources
- [Install Ansible OSERROR Python Fix](https://stackoverflow.com/questions/54778630/could-not-install-packages-due-to-an-environmenterror-errno-2-no-such-file-or)
- [Course Catalog](https://myhub.njit.edu/BannerExtensibility/customPage/page/stuRegCrseSched)
- [Golang: Build from source](https://go.dev/doc/install)
- [MySQL New User](https://askubuntu.com/questions/1322175/not-allowed-to-create-user-with-grant)
- [Avoid SQL Injection](https://go.dev/doc/database/sql-injection)
- [nginx proxy_pass server set_up](https://serverfault.com/questions/598202/make-nginx-to-pass-hostname-of-the-upstream-when-reverseproxying)
- [OG Go Digitalocean Service](https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-using-nginx-on-ubuntu-18-04#conclusion)
- [Lets encrypt nginx: ubuntu 20.04](https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-ubuntu-20-04)