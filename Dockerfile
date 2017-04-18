FROM alpine:3.5

ADD dhcpd-monitor /dhcpd-monitor
ADD dhcpd-pools /bin/dhcpd-pools

ENV PORT 80
ENV DHCPD_CONF_FILE /etc/dhcpd.conf
ENV DHCPD_RESERVATION_FILE /etc/dhcpd-reservation.conf
ENV DHCPD_LEASE_FILE /etc/dhcpd.leases

ENTRYPOINT ["/dhcpd-monitor"]

