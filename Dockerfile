FROM golang

COPY . /go/src/git.klimlive.de/frzifus/dbwt

#RUN ln -s /go/src/git.klimlive.de/frzifus/dbwt/view view
#RUN ln -s /go/src/git.klimlive.de/frzifus/dbwt/static static
#RUN ln -S /go/src/git.klimlive.de/frzifus/dbwt/config config

COPY ./view /go/bin/view
COPY ./static /go/bin/static
COPY ./config /go/bin/config

WorkDir /go/src/git.klimlive.de/frzifus/dbwt

RUN go get -v

RUN go install

WorkDir /go/bin/

ADD ./init/wait-for-it.sh /go/bin/

RUN chmod +x wait-for-it.sh

EXPOSE 8090

CMD ["./wait-for-it.sh", "database:3306", "--", "dbwt"]
