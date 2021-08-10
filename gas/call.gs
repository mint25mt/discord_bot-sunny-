const heroku_url = 'https://coffe-break-discordchan.herokuapp.com/';
const s = new Date(); s.setDate(s.getDate() - 1);
const now = s.getFullYear() + "/" + (s.getMonth() + 1) + "/" + (s.getDate());


function discordCorona() {
  Logger.log(now);
  UrlFetchApp.fetch(heroku_url + "?content=corona&date=" + now);
}

function discordWeather() {
  UrlFetchApp.fetch(heroku_url + "?content=weather");
}