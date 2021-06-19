ARG KRAKEND_VERSION
FROM devopsfaith/krakend:${KRAKEND_VERSION}

LABEL maintainer="phuc.ps.89@gmail.com"

# RUN useradd -r -c "KrakenD user" -U krakend

USER krakend

VOLUME [ "/etc/krakend" ]

WORKDIR /etc/krakend

# ENTRYPOINT [ "/usr/bin/krakend" ]

# CMD [ "run", "-c", "/etc/krakend/krakend.json" ]

# EXPOSE 8080 8090