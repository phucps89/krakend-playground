FROM graylog2/server:2.1.1-1
COPY udp-input-graylog.json /usr/share/graylog/data/contentpacks
RUN chown -Rf graylog:graylog /usr/share/graylog/data/contentpacks/udp-input-graylog.json
