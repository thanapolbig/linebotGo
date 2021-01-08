# office-new-normal-golang

GOLANG :
IMPORT : https://gitlab.com/ohmsm0150/office-new-normal

# LineBot

โปรเจ็ค linebot นี้มีการใช้ภาษา golang ในการพัฒนา โดยมีการต่อ database ที่เป็นของ sqlserver ที่อยู่บนเครื่อง clound AWS 

## lineBot
ก่อนที่จะทำการเขียน go เราต้องทำการสมัคร [linedev](https://developers.line.biz/en/) เสียก่อน โดยพอสมัรแล้วให้ทำการสร้าง bot ของเรา โดยเลือกประเภทเป็น Messaging API กรอกรายละเอียดให้เรียบร้อย โดยพอสมัครแล้วจะมีส่วนสำคัญต่างๆที่ต้องใช้ในการ dev คือ

* **Webhook** ตั้งค่าเป็น Enable และช่องใส่ link ที่เราต้องเอา link server ที่เราทำเสร็จแล้วมาใส่ไว้

* **Channel Secret** ที่เราต้องนำจาก line bot ไปใส่ใน server ของเรา เพื่อระบุถึงตัวตน Line bot ID ที่เราต้องการเชื่อมต่อ

* **Chanel Token** คือ token ที่จะใช้ในการเชื่อมต่อกับ line bot ของเรา

## Installation Golang

สามารถเข้าไปโหลดภาษา go ได้ที่ [golang](https://golang.org/doc/install)

1. คำสั่งโหลด go 

```bash
pip install foobar
```
2. ทำการ set PATH environment ต่างๆ 

```bash
export PATH=$PATH:/usr/local/go/bin
```
3. ก็ทำการเช็ค version ของ go เราได้เลย
```bash
go version
```
## Installation Ngrok
ngrok เป็นเครื่องมือที่ทำให้ web หรือ api ที่ run บน localhost สามารถ online ได้นั่นเอง ในที่นี้เราเอามาแปลง url ของ api เราให้ online  และทำเป็น https

[install ngrok](https://ngrok.com/download)

## How to use
คำสั่ง run api 
```
go run .
```
พอ run api เสร็จให้ทำการ run ngrok โดย
```
ngrok [เลข port ที่run api]
```
หลังจากนั้น ก็จำได้ url ที่ online แล้วเป็นแบบ http และ https โดย Webhook ของlinebot ต้องใช้เป็น https 

## License
[medium](https://medium.com/@khemcharoenreadyma/%E0%B8%A1%E0%B8%B2%E0%B8%97%E0%B8%B3-line-bot-%E0%B9%84%E0%B8%A7%E0%B9%89%E0%B9%80%E0%B8%8A%E0%B9%87%E0%B8%84-web-service-%E0%B9%80%E0%B8%A5%E0%B9%88%E0%B8%99%E0%B8%81%E0%B8%B1%E0%B8%99%E0%B9%80%E0%B8%96%E0%B8%AD%E0%B8%B0-by-golang-90f99b9fa56f)
