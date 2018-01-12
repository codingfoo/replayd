replayd:
  user.present:
    - fullname:
    - shell: /bin/nologin
    - home: /usr/local/replayd

/usr/local/bin/replayd:
  file.managed:
    - source: salt://replayd
    - mode: 755

/etc/replayd:
  file.directory:
    - dir_mode: 755

/etc/replayd/config.json:
  file.managed:
    - source: salt://config.json
    - mode: 644

/etc/systemd/system/replayd.service:
  file.managed:
    - source: salt://replayd.service
    - mode: 755

replayd-service:
  service.running:
    - name: replayd
    - enable: true
    - watch:
      - file: /usr/local/bin/replayd
      - file: /etc/replayd/config.json
      - file: /etc/systemd/system/replayd.service
