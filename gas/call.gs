const heroku_url = 'https://coffe-break-discordchan.herokuapp.com/';

function discordCorona() {
  UrlFetchApp.fetch(heroku_url + "?=display=corona");
}

function discordWeather() {
  UrlFetchApp.fetch(heroku_url + "?=display=weather");
}
