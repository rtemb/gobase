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

Add variables DOCKER_EMAIL, DOCKER_USER, DOCKER_PASS varibles in _your_travis_ci_repo_/settings

### 5. Deploy in Heroku
Instal Heroku https://devcenter.heroku.com/articles/heroku-cli

For first time create and deploy you app manualy:
```
heroku login
heroku apps:create gobase
heroku plugins:install heroku-container-registry
heroku container:login
heroku container:push web
heroku ps:scale web=1 
```

Add HEROKU_API_KEY, HEROKU_APP_NAME varibles in Travis CI page _your_travis_ci_repo_/settings
HEROKU_API_KEY - Heroku API can be found in ‘Account Settings’, in heroku (https://dashboard.heroku.com/account)
HEROKU_APP_NAME - it's result of command _heroku create_

