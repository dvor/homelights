# Homelights

Homelights is a small app that I use to control Philips Hue lights at my apartment.

I have two budgerigars which require specific light schedule. Birds are awake from 8:00 till 20:00 in winter, and from 8:00 till 22:00 in summer.

App starts a daemon that toggles lights on/off based on local time, season and weather conditions (there might be not enough light outside because of the clouds, etc.). Weather data is fetched from openweathermap.org.

# Dependencies

Dependencies are managed by dep. See [Daily Dep](https://golang.github.io/dep/docs/daily-dep.html).
