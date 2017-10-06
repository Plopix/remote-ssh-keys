# Remote SSH Keys

Simple tool that use a YAML file to lookup for SSH Keys.

# Install

Once you have the binary, you just have to configure your SSH server to use it:

```bash
mv remote-ssh-keys /usr/bin
chmod 700 /usr/bin/remote-ssh-keys
chown root:root /usr/bin/remote-ssh-keys
touch /etc/remote-ssh-keys.conf
```

Tell you SSH server

```
# /etc/ssh/sshd_config
AuthorizedKeysCommand /usr/bin/remote-ssh-keys %u %h %t %f %k
AuthorizedKeysCommandUser root
```

Then you need to configure the YAML file

Ex:

```yaml
# /etc/remote-ssh-keys.conf
---
accounts:
    s.morel:
        github:
            - "plopix"
        url:
            - "https://somehere.com/Plopix.keys?%s"
        plain:
            - "plaintext_plopix"
```

Here that is pretty simple, if someone try to login as `s.morel`, the system
will loop and look online on each providers (github, url, and plain) to find
public keys. If the SSH client match a public key with a private proposed then
the user is logged.

This little tool avoid you to manage your SSH Keys.

```bash
/etc/init.d/ssh restart
```
