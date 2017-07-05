Tamplate for Go app

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

Run container with <internal port>:<exernal port>: 
```
docker run -e PORT='8081' -p 8081:8081 gobase
```