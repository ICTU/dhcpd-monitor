# DHCPD monitor

## Introduction
This service wraps [dhcpd-pools](http://dhcpd-pools.sourceforge.net/man.html) and exposes a web endpoint to retrieve state of an isc-dhcp-server instance.

```
curl localhost/v1/api/state
```

returns

```
{
  "subnets": [
    {
      "location": "All networks",
      "range": "10.25.166.21 - 10.25.166.40",
      "defined": 20,
      "used": 0,
      "touched": 0,
      "free": 20
    },
    {
      "location": "All networks",
      "range": "10.25.166.41 - 10.25.166.50",
      "defined": 10,
      "used": 0,
      "touched": 0,
      "free": 10
    },
    {
      "location": "All networks",
      "range": "10.25.166.51 - 10.25.166.240",
      "defined": 190,
      "used": 0,
      "touched": 1,
      "free": 190
    }
  ],
  "shared-networks": [],
  "summary": {
    "location": "All networks",
    "defined": 220,
    "used": 0,
    "touched": 1,
    "free": 220
  }
}

```

## How to run

```
PORT=12345 DHCPD_CONF_FILE=/tmp/dhcpd.conf DHCPD_LEASE_FILE=/tmp/dhcpd.leases docker run --rm -it -p 12345:12345 -v /tmp/dhcpd.conf:/tmp/dhcpd.conf -v /tmp/dhcpd.leases:/tmp/dhcpd.leases -e PORT=12345 -e DHCPD_CONF_FILE=/tmp/dhcpd.conf -e DHCPD_LEASE_FILE=/tmp/dhcpd.leases ictu/dhcpd-monitor
```
