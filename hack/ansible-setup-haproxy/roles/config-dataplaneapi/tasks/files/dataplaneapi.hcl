# https://github.com/haproxytech/dataplaneapi/tree/master/configuration

config_version = 2
name = "haproxy1"
mode = "single"

dataplaneapi {
  scheme = ["http"]
  host = "0.0.0.0"
  port = "5555"
  tls {
    tls_host = "null"
    tls_port = 6443
    tls_certificate = "null"
    tls_key = "null"
    tls_ca = "null"
  }
  userlist {
    userlist = "controller"
    userlist_file = "/etc/haproxy/userlist.cfg"
  }
  transaction {
    transaction_dir = "/tmp/haproxy"
  }
}

haproxy {
  config_file = "/etc/haproxy/haproxy.cfg"
  haproxy_bin = "/usr/sbin/haproxy"
  reload {
    reload_delay = 5
    reload_cmd = "systemctl reload haproxy"
    restart_cmd = "systemctl restart haproxy"
    status_cmd = "systemctl status haproxy"
    service_name = "haproxy.service"
  }
}

log_targets = [
  {
    log_to           = "stdout"
    log_level        = "debug"
    log_format       = "json"
    log_types = [
      "access",
      "app",
    ]
  },
  {
    log_to           = "file"
    log_file         = "/var/log/dataplanepi.log"
    log_level        = "info"
    log_format       = "text"
    log_types        = ["app"]
  },
  {
    log_to           = "syslog"
    log_level        = "info"
    syslog_address   = "127.0.0.1"
    syslog_protocol  = "tcp"
    syslog_tag       = "dataplaneapi"
    syslog_level     = "debug"
    syslog_facillity = "local0"
    log_types        = ["access"]
  },
]
