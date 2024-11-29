
# TenzInformer

The project is aimed at working with collecting information and sending it to tenzir.

## Installation
```bash
  git clone https://github.com/kudras3r/TenzInformer.git
  cd TenzInformer
  chmod +x informer.sh 
  sudo ./informer.sh install 
```

## Unstallation
In TenzirInformer dir:
```bash
  sudo ./informer.sh uninstall
```


## Usage/Examples

```bash
  sudo systemctl enable tenzinformer.service
  sudo systemctl start tenzinformer.service
  sudo systemctl status tenzinformer.service
  sudo systemctl stop tenzinformer.service
  sudo systemctl disable tenzinformer.service
```

## Roadmap

- 25/11/24 | Initial commit.
- 25/11/24 | Init logger.
- 26/11/24 | Add grabber.
- 27/11/24 | Add sender. Final app building steps.
- 28/11/24 | Add installer / uninstaller.


## Authors

- [@kudras3r](https://www.github.com/kudras3r)

