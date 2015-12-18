# OpenVPN NetworkManager Split

Because NetworkManager can't read the new way of OpenVPN to create client config files.

This little tool will split the file.ovpn into 4 files (.ca.pem, .cert.pem, .key.pem, .tls-auth.pem) so you can directly link them in your NetworkManager configuration UI.


## How to use

Download your .ovpn file, then

```
ovpnsplit file.ovpn
```

and you should have

```
file.ovpn.ca.pem
file.ovpn.cert.pem
file.ovpn.key.pem
file.ovpn.tls-auth.pem
```

### Configure NetworkManager

#### Add a new network
![add-net](doc/01-add-net.png)

#### Import a .ovpn file
![import-vpn-config](doc/02-import-vpn-config.png)

#### Select the file
![select-ovpn-file](doc/03-select-ovpn-file.png)

#### Open the advanced config panel
![advanced-config](doc/10-advanced-config.png)

#### Go to TLS Auth Tab
![tls-auth](doc/11-tls-auth.png)

#### Set the key direction
![set-key-direction](doc/15-set-key-direction.png)

## Limitations

Right now, it is based on .ovpn files created by the [kylemanna/openvpn](https://hub.docker.com/r/kylemanna/openvpn) docker container, which handle tls-auth and stuff. Your .ovpn file might differ and you might not get all those files.
