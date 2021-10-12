FROM golang:latest
# directory for the apps source 
RUN mkdir /app

# copy everything from the root dir into /app
ADD . /app

# specify workdir location 
WORKDIR /app

# download any go deps
RUN go mod download

# run go build to create a binary to execute.
RUN go build -o main . 

# command to start off the new app
CMD ["/app/main"]