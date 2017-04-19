FROM alpine:3.5

ADD dhcpd-monitor /dhcpd-monitor
ADD dhcpd-pools /bin/dhcpd-pools

ENV PORT 80
ENV DHCPD_CONF_FILE /etc/dhcpd.conf
ENV DHCPD_LEASE_FILE /etc/dhcpd.leases

CMD /dhcpd-monitor --port=$PORT --dhcpd-config-file=$DHCPD_CONF_FILE --dhcpd-lease-file=$DHCPD_LEASE_FILE
