FROM golang:1.18
WORKDIR /app
COPY . .
RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air
EXPOSE 8888
CMD ["/bin/air"]
