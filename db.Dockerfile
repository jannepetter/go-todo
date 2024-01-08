FROM mongo:7.0

EXPOSE 27017

VOLUME /app/data

CMD ["mongod"]