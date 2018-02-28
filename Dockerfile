FROM partlab/ubuntu-golang

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone
RUN dpkg-reconfigure -f noninteractive tzdata

RUN mkdir /src
ADD . /src
RUN go get github.com/gin-gonic/gin
RUN go get github.com/garyburd/redigo/redis
RUN go get gopkg.in/yaml.v2

WORKDIR /src
RUN chmod -R 777 /src

