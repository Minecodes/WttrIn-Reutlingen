import requests
import sys
args = sys.argv

# Request weather data from wttr.in
def get_weather(location, language):
    url = 'http://wttr.in/' + location + '?0&T&lang=' + language
    res = requests.get(url, headers={'User-Agent': 'curl/7.58.0'})
    return res.text

def post_weather(misskey, token, text):
    res = requests.post(misskey + '/api/notes/create', json={'i': token, 'text': text, 'visibility': 'public', 'localOnly': False},
                  headers={'User-Agent': 'wttrinReutlingen Bot (v2)', 'Authorization': token})
    print(res.text)

def main():
    if len(args) < 3:
        print('Usage: python bot.py <misskey url> <token>')
        return
    instance = args[1]
    token = args[2]
    city = "Reutlingen"
    language = "de"

    # Download the weather forecast
    weather = get_weather(city, language)

    print(weather)
    # Post the weather forecast to Misskey
    post_weather(instance, token, weather)

main()