# dumb
With the massive daily increase of useless scripts on Genius's web frontend, and having to download megabytes of clutter, [dumb](https://github.com/rramiachraf/dumb) tries to make reading lyrics from Genius a pleasant experience, and as lightweight as possible.

<a href="https://codeberg.org/rramiachraf/dumb"><img src="https://img.shields.io/badge/Codeberg-%232185d0" /></a>

![Screenshot](https://raw.githubusercontent.com/rramiachraf/dumb/main/screenshot.png)

## Installation & Usage
### Docker
```bash
docker run -p 8080:5555 --name dumb ghcr.io/rramiachraf/dumb:latest
```

### Build from source
[Go 1.20+](https://go.dev/dl) is required.
```bash
git clone https://github.com/rramiachraf/dumb
cd dumb
make build
./dumb
```

The default port is 5555, you can use other ports by setting the `PORT` environment variable.

## Public Instances

| URL                                           | Tor                                                                               | I2P                                                                        | Region | CDN? | Operator               
| ---                                           | ---                                                                               | ---                                                                        | ---    | ---  | ---                    
| <https://dumb.ducks.party>                    | No                                                                                | No                                                                         | NL     | No   | https://ducks.party 
| <https://dm.vern.cc>                          | [Yes](http://dm.vernccvbvyi5qhfzyqengccj7lkove6bjot2xhh5kajhwvidqafczrad.onion)   | [Yes](http://vernxpcpqi2y4uhu7to4rnjmyjjgzh3x3qxyzpmkhykefchkmleq.b32.i2p) | US     | No   | https://vern.cc        
| <https://dumb.lunar.icu>                      | No                                                                                | No                                                                         | DE     | Yes  | @MaximilianGT500       
| <https://dumb.privacydev.net>                 | [Yes](http://dumb.g4c3eya4clenolymqbpgwz3q3tawoxw56yhzk4vugqrl6dtu3ejvhjid.onion) | No                                                                         | FR     | No   | https://privacydev.net 
| <https://dumb.privacyfucking.rocks>           | No                                                                                | No                                                                         | DE     | -    | https://privacyfucking.rocks |
| <https://sing.whatever.social>                | No                                                                                | No                                                                         | US/DE  | Yes  | Whatever Social        

#### Notes:
- Instances list in JSON format can be found in [instances.json](instances.json) file.
- For people who might be capable and interested in hosting a public instance feel free to do so, and don't forget to open a pull request, so your instance can be included here.

## Contributing
Contributions are welcome.

## License
[MIT](https://github.com/rramiachraf/dumb/blob/main/LICENCE)

