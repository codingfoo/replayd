replayd:
  user.present:
    - shell: /usr/sbin/nologin
    - home: /dev/null

/etc/replayd:
  file.directory:
    - dir_mode: 755

go-install:
  pkg.installed:
    - pkgs:
      - golang

replayd-tarball:
  file.managed:
    - name: /tmp/replayd-v1.1.tar.gz
    - source: https://github.com/codingfoo/replayd/archive/v1.1.tar.gz
    - source_hash: md5=b8ca32713f4c3dbd5e178b068376ca4d

extract-replayd:
  cmd:
    - cwd: /tmp
    - names:
      - tar xvf replayd-v1.1.tar.gz
    - run
    - require:
      - file: replayd-tarball

build-replayd:
  cmd:
    - cwd: /tmp/replayd-1.1
    - names:
      - go build main.go
    - run
    - require:
      - cmd: extract-replayd

/usr/local/bin/replayd:
  file.managed:
    - source: /tmp/replayd-1.1/main
    - mode: 755
    - require:
      - cmd: build-replayd
      - file: /etc/replayd/config.json

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
