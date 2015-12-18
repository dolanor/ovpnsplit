# Ovpnsplit

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


Add a new network
![add-net](doc/01-add-net.png)

Import a .ovpn file
![import-vpn-config](doc/02-import-vpn-config.png)

Select the file
![select-ovpn-file](doc/03-select-ovpn-file.png)

Choose a user cert file
![click-user-cert](doc/04-click-user-cert.png)

Select the file
![select-user-cert](doc/05-select-user-cert.png)

Choose a CA cert file
![click-ca-cert](doc/06-click-ca-cert.png)

Select the file
![select-ca-cert](doc/07-select-ca-cert.png)

Choose a key file
![select-private-key](doc/08-select-private-key.png)

Select the file
![select-private-key](doc/09-select-private-key.png)

Open the advanced config panel
![advanced-config](doc/10-advanced-config.png)

Go to TLS Auth Tab
![tls-auth](doc/11-tls-auth.png)

Add a optional TLS authentication
![add-tls-auth-file](doc/12-add-tls-auth-file.png)

Choose a tls-auth file
![choose-tls-auth-file](doc/13-choose-tls-auth-file.png)

Select the file
![select-tls-auth-file](doc/14-select-tls-auth-file.png)

Set the key direction
![set-key-direction](doc/15-set-key-direction.png)

## Limitations

Right now, it is based on .ovpn files created by the [kylemanna/openvpn](https://hub.docker.com/r/kylemanna/openvpn) docker container, which handle tls-auth and stuff. Your .ovpn file might differ and you might not get all those files.
