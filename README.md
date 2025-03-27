# go-weather
It is a simple weather CLI integrated with weather API https://wttr.in/

# Usage
you can simply start this app
```cmd
go run .
```
# Flags:
`--city`

### Description:
Specifies the city for which you want to retrieve weather information. 

### Default value
Moscow

### Usage example
```cmd
go run . --city=Paris
```
Output:
```cmd
Cheking weather
☀️   +16°C
```


`--format`

### Description:

Specifies the weather format.

1 = `☀️   +16°C`

2 = `☀️   🌡️+16°C 🌬️↙10km/h`

3 = `Paris: ☀️   +16°C`

4 = `Paris: ☀️   🌡️+16°C 🌬️↙10km/h`
### Default value
1 = `☀️   +16°C`
### Usage example

```cmd
go run . --city=Paris --format=4
```

Output:

```cmd
Cheking weather
Paris: ☀️   🌡️+16°C 🌬️↙10km/h
```