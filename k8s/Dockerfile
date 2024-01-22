FROM envoyproxy/envoy:v1.24.0

# override the existing default envoy.yaml 
COPY envoy.yaml etc/envoy/envoy.yaml

RUN chmod go+r /etc/envoy/envoy.yaml
COPY ./helloworld.proto /etc/envoy/protos/
COPY ./helloworld.pb /etc/envoy/protos/
COPY ./book.proto /etc/envoy/protos/
COPY ./book.pb /etc/envoy/protos/


CMD ["/usr/local/bin/envoy", "-c", "/etc/envoy/envoy.yaml"]