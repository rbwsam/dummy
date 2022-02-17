# syntax=docker/dockerfile:1
FROM scratch
COPY bin/dummy /
CMD ["/dummy"]