FROM ubuntu
ENV LC_ALL en_US.UTF-8
RUN apt-get update 
RUN echo 8 | apt-get install -y tzdata
COPY ./app /tmp/linxdatacenter
COPY ./app/data /var/data
RUN apt-get install -y golang && \
 cd /tmp/linxdatacenter && go build -o /bin/app
#RUN /etc/init.d/php7.2-fpm start
#RUN /etc/init.d/nginx restart
#ADD ~/www/netselect /var/www/html
#CMD ["nginx", "-g", "daemon off;"]
CMD  
ENTRYPOINT ["/bin/app"]
CMD ["/var/data/1.json"]