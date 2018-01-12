replayd:
  user.present:
    - fullname:
    - shell: /bin/nologin
    - home: /usr/local/replayd

go-install:
  pkg.installed:
    - pkgs:
      - golang

replayd-repo:
  git.latest:
    - name: git@github.com:codingfoo/replayd.git
    - branch: stable
    - target: /tmp/replayd

/usr/local/bin/replayd:
  file.managed:
    - source: /tmp/replayd/replayd
    - mode: 755

/etc/replayd:
  file.directory:
    - dir_mode: 755

/etc/replayd/config.json:
  file.managed:
    - mode: 644
    - contents: |
      {
      "host": "127.0.0.1",
      "port": "8080"
      }

/etc/systemd/system/replayd.service:
  file.managed:
    - mode: 755
    - contents: |
      [Unit]
      Description=HTTP replayd script
      After=network.target

      [Service]
      Type=simple
      ExecStart=/usr/local/bin/replayd -config-file=/etc/replayd/config.json
      Restart=always
      User=replayd
      Group=replayd

      [Install]
      WantedBy=multi-user.target

replayd-service:
  service.running:
    - name: replayd
    - enable: true
    - watch:
      - file: /usr/local/bin/replayd
      - file: /etc/replayd/config.json
      - file: /etc/systemd/system/replayd.service
