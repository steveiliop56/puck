## Puck 🏒

Puck (Package Update Checking Kit _yes I know it sucks_) is a simple tool that connects to your servers and checks for
apt package updates.

<img alt="Screenshot" src="screenshots/screenshot.png" width="545" height="299">

> Warning ⚠️: Puck is in early stages of development it still has a lot of missing features I want to add. I hope I can have a stable release soon.

### Features 😲

Well about that...

- Tiny CLI (only 10mb)
- Beautiful CLI
- Really fast\*

> Puck itself is really fast since it's written in GO but the actual speed depends on the server

### Todo 📃

- [x] Redesign the CLI UI
- [ ] Docker Image
- [ ] Web UI
- [ ] Notifications (via ntfy/discord)
- [x] Ability not to use sudo (for systems running with root)
- [ ] Update systems?
- [x] Support for other package managers (currently only supporting apt)

### Running 🏃

Puck is built for multiple architectures and systems and you can simply download it and run it from the release page, thats it!

### Building 🛠️

You can build `puck` very easily by installing go and git on your pc and then simply running:

```bash
go build .
```

The build is really fast and when it finished you should have a binary called `puck` in your current directory.

### Usage

Puck is really simple to use, it works using a simple yaml configuration file. Here is an example:

```yaml
servers:
  - name: myserver
    hostname: 192.168.1.5
    username: someone
    password: hello!

  - name: server2
    hostname: server2.local
    username: me
    password: reallysecurepassword
    privateKey: /some/path/id_rsa
    noSudo: true
```

Here is the reference table for the available options:

| Name         | Description                                       | Type      | Required |
| ------------ | ------------------------------------------------- | --------- | -------- |
| `name`       | Name of the server you can put whatever you want. | `string`  | yes      |
| `hostname`   | IP or hostname of the server.                     | `string`  | yes      |
| `username`   | Username for ssh.                                 | `string`  | yes      |
| `password`   | Password used for ssh and for sudo.               | `string`  | yes      |
| `privateKey` | Private key path to use for ssh.                  | `string`  | yes      |
| `noSudo`     | Don't use sudo to run the commands.               | `boolean` | no       |

After you make your configuration file you can use puck like so:

```bash
./puck check
```

Puck be default will look for `puck.yml` but if you wish to use a different file name you can use the `-c` or `--config` flag to specify a custom path.

### Contributing ❤️

Contributing is really easy in puck you simply need to have go and git in your system, then you can clone the repository make your changes and open a pull request. Any help is appreciated.

### License 📜

The project is licensed under the GPL V3 License. You may modify, distribute and copy the code as long as you keep the changes in the source files. Any modifications you make using a compiler must be also licensed under the GPL license and include build and install instructions.

### Acknowledgments 🙏

- The project is heavily inspired from [cup](https://github.com/sergi0g/cup).
- [Carbon](https://carbon.now.sh/) thanks for the screenshot
