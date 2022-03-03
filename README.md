# empty-room :construction: `Under Construction`
Infra + Dev into service that queries empty rooms @ NJIT

## Site
- Site will be hosted here: "" 

## :tram: Infrastructure
```txt
.------------.             .-------------.
| Front-end  |   <----->   | API Service |
'------------'             '-------------'
```

## :hammer: How to Build
- This repo encapsulates multiple components needed for the program to work
- Ensure that path in `src/api/server.go` corresponds to out.json

## :microscope: Technologies
- Languages: `go`, `typescript`, `css`, `html`
- Frameworks: `react`
- Packages: `nginx`, `ngrok`, `npm`

## Infra
- [Golang + Nginx Setup](https://github.com/hahdookin/cs490/blob/main/util/dep.sh) 
- [Infrastructure Confs](https://github.com/hahdookin/cs490/tree/main/infra)


## :microscope: Directory Explanation

| Directory         | Explanation
| :---:             | :---
| [src/front-end](src/front-end)            | ReactJS Front-end
| [src/api](src/api) | Endpoint to get information about empty-rooms @ NJIT
| [data-processing](data-processing)| Data Processing from data-collection dir

## :books: Resources
- [Install Ansible OSERROR Python Fix](https://stackoverflow.com/questions/54778630/could-not-install-packages-due-to-an-environmenterror-errno-2-no-such-file-or)
- [Course Catalog](https://myhub.njit.edu/BannerExtensibility/customPage/page/stuRegCrseSched)
- [Golang: Build from source](https://go.dev/doc/install)
- [MySQL New User](https://askubuntu.com/questions/1322175/not-allowed-to-create-user-with-grant)
- [Avoid SQL Injection](https://go.dev/doc/database/sql-injection)
- [React: Creating a new react app](https://reactjs.org/docs/create-a-new-react-app.html)
- [nginx proxy_pass server set_up](https://serverfault.com/questions/598202/make-nginx-to-pass-hostname-of-the-upstream-when-reverseproxying)
