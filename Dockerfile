FROM registry.access.redhat.com/ubi8-minimal

EXPOSE 8442

RUN microdnf update -y && rm -rf /var/cache/yum && microdnf install git go make -y && microdnf clean all

COPY . /opt/golang-demo
WORKDIR /opt/golang-demo

RUN make build

CMD ["/opt/golang-demo/bin/golang-demo"]

