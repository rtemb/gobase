Tamplate for Go app

### Travis CI Status: [![Build Status](https://travis-ci.org/rtemb/gobase.svg?branch=master)](https://travis-ci.org/rtemb/gobase)

### 1. Install glide (https://github.com/bumptech/glide)
```
curl https://glide.sh/get | sh
```

To install dependencies run:
```
glide install
```


### 2. Install Yelp (http://pre-commit.com)
```
curl https://bootstrap.pypa.io/get-pip.py | sudo python - pre-commit
```

Install pre-commit hook:
```
pip install pre-commit
```

To check all files: 
```
pre-commit run --all-files
```

### 3. Build app 
Build for linux:
```
GOOS=linux CGO_ENABLED=0 go build
```

### 4. Docker

Build docker container: 
```
docker build -t gobase -f ./Dockerfile .
```

Run container with env var PORT=8081 and _exernal port_:_internal port_: 
```
docker run -e PORT='8081' -p 8081:8081 gobase
```

Create DockerHub repository for CI: https://hub.docker.com 

### 5. Travis CI 

Enable Travis CI for repository here: https://travis-ci.org/profile
Click "Sync account" and enable CI for your repository

Add DOCKER_EMAIL, DOCKER_USER, DOCKER_PASS varibles in _your_travis_ci_repo_/settings

