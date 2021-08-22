# Quack

Simple service in Go to keep you IP updated on Duck DNS

## Instruction

1. Register your domain on https://www.duckdns.org/
2. create a configuration file named `config.json` and put it in the same folder where the app is installed.

You can change set a custom path for your configuration file with the `QUACK_CONFIG` environment variable.

### Configuration File

```json
{
  "token": "YOUR_TOKEN",
  "domain": "YOUR_DOMAIN"
}
```