## Komponentit
#### Nettisivu frontend
###### Techstack
- Nuxt.js?
- Tailwind
- Authentikaatio aromiii auth palvelulta
#### Puhelin frontend
###### Kuvaus
- Pystyy kuunnella musiikkia joka striimataan backendistä
- Pystyy lataamaan musiikkia netittömään soittoon
- Kaikki tärkeä ominaisuudet
	- Uniajastin
	- Taustalla soitto
	- Ei bugita kuten yt music
###### Techstack
- Flutter
#### Authentikaatio palvelu
###### Kuvaus
- Samalla authentikaatio front ja backend. 
- Käytää next-authia johon api endpoitteihin tekee paremmat wrapperit. 
- account.aromiii.com hostattu jossa voi hallita aromiii ecosysteemin käyttäjää
	- Google tyylinen ja käytännössä kopio omalla designillä
	- Livealiin pitää vaihtaa kun tämä valmis
- Googlen tyylinen session management ja javascript component jolla yläkulmassa ikkuna jossa voi vaihtaa helposti käyttäjiä yms
###### Techstack
- Next.js
- Next-auth
- Tailwind
#### Audio player backend
###### Kuvaus
- Striimaa audiofilen palasina frontendiin jotta musiikin soittaminen alkaa lähes heti
	- Samalla tavalla kun spotify, yt music yms
- Pystyy latamaan serverille omia biisejään ja kuunnella niitä netin yli
- Voi ladata serveriltä musiikkia locaaliksi jotta kun ei pääse nettiin toimii
- Pystyy converttaa spotify ja yt music soittolistan
	- Siirtyminen olisi mahdollisimman helppoa
- Älykäs shuffle soittolistoihin
###### Techstack
- Go
- Websockets striimaukseen