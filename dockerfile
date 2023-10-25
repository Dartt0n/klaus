FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/klaus . 


FROM scratch
COPY --from=0 /bin/klaus /bin/klaus
CMD ["/bin/klaus"]