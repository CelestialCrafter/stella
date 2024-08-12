FROM node:22.6-alpine as web
WORKDIR /app/server/web
COPY server/web/package.json .
COPY server/web/package-lock.json .
RUN npm i

COPY server/web .
RUN npm run build

FROM golang:1.22-alpine
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
COPY --from=web . server/web
RUN apk add g++ blender-headless py3-numpy
RUN CGO_ENABLED=1 GOOS=linux go build -o /stella

EXPOSE 80
ENV ADDRESS :80
ENV BLENDER_EXE blender-headless
ENV BLENDER_DATA_PATH /app/blender/
CMD ["/stella"]
