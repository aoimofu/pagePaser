# pagepaser
It's html parse command for website.

## Description
With this command you can retrieve the necessary data from the website.

## Usage
```
pagepaser -u URL -e HTML ELEMENT
```

## Example
```
pagepaser -u http://www.lovelive-anime.jp/uranohoshi/news.php -e 'div.title'

Output:
Hybrid Mind Market×『ラブライブ！サンシャイン!!』セレクトショップ　　新商品のお知らせ
ＰＬＥＸより新商品のお知らせ
3月6日発売「ダ・ヴィンチ4月号」にて、『ラブライブ！サンシャイン!!』の20P超の大特集が決定！
『ラブライブ！サンシャイン!!』もっと輝け!! Aqours ３号連続カバーガール総選挙PART1～TVアニメ２期衣装～、中間発表！
ダイビング専門誌『DIVER』×ラブライブ！サンシャイン!!　松浦果南描きおろしグッズ販売のお知らせ
```
