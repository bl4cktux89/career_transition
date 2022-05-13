FROM python:alpine3.15 AS build

WORKDIR /mkdocs
COPY ./ /mkdocs
RUN pip install -U mkdocs
RUN mkdocs build

FROM nginx:stable-alpine AS release
COPY --from=build /mkdocs/site /usr/share/nginx/html