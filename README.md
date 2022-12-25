# demo-1
demo-1 for Dashboard.

## Öncelikle bigisayarınıza yüklmeniz gereken dosyalar

1. https://nodejs.org/en/  adresinden nodejsi indirip bilgisayarınza kurunuz. <br/>
2. Powershelli Administrator olarak çalıştırnın ve `npm init vue@latest` komutu girerek vuejsi bilgisayariniza kurun.


## Sistemi ayağa kaldırmadan önce yapılması gereken adımlar:

### 1. Database işlemleri:

Öncellikle `demo-1` klasorünün altındaki db dosyalarının bilgileri ile `postgres` üzerinde `chartdatas` ve `register` isimlerine sahip olan iki adet database oluşturun. <br/>
Postgresqlde oluşturduğunuz `chartdatas` databaseıne `query tool` açın ve sol üstte bulunan open file kısmından demo-1/db klasorünün bulunduğu pathe gidin `chartdatas.db` dosyaini seçin ve open butonuna basın ve `run` butonuna basin. <br/>
Ayni islemi `register` içinde uygulayın.

### 2. go işlemleri

`main.go` satir `39` ve `register.go` satir `20` de bulunan password degiskenine `postgres` şifrenizi giriniz.

### 3. vue js işlemleri

Terminalde `demo-1` dizininde iken aşağıdaki kodları çalıstırn: <br/>

`yarn install` <br/>
`npm install chart.js`

### 4. Aşağıdakı komutlar ile sistemi ayağa kaldırnız

Sistemi ayağa kaldıramk için bir adet go için ve bir adet vuejs için iki adet terminal lazımdır. <br/> <br/>

İlk terminal ile `demo-1/go-backend` dosyasina gidiliniz ve burda terminale `go run main.go` yazınız. <br/>
İkinci terminal  `demo-1` dizininde `yarn serve` komutu çalıştırınız. <br/>



https://user-images.githubusercontent.com/72500821/209441383-6289c64b-ca33-45b1-a0e8-3f3ab37b6591.mp4


