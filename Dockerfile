FROM golang:1.13
COPY ./app /tmp/linxdatacenter
COPY ./app/data /var/data
RUN cd /tmp/linxdatacenter && go build -o /bin/app
ENTRYPOINT ["/bin/app"]
CMD ["/var/data/1.json"]