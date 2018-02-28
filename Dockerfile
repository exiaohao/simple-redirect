FROM partlab/ubuntu-golang
RUN mkdir /src
ADD . /src
RUN go get github.com/gin-gonic/gin
RUN go get github.com/garyburd/redigo/redis
RUN go get gopkg.in/yaml.v2

WORKDIR /src
RUN chmod -R 777 /src