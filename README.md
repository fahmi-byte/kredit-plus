This application isn't complete yet!

To run this application follow these step:

- first install all module for this apps
- Insert table merchant masih manual dengan cara menjalankan file test/create_environment_test.go cari fungsi untuk insert atau bisa manual melalui DBMS
- Jika sudah bisa melakukan registrasi melalui collection Postman
- Lalu jika akun sudah terbuat, lanjut login dengan credential yang sudah dibuat sebelumnya
- Selesai login akan mendapatkan token jwt yang nantinya akan terpakai untuk authorization API
- Gunakan token di semua API kecuali register, login, dan /installment-process
- Khusus untuk API installment-process memiliki apiKey tersendiri yang dimiliki tiap merchant untuk otorisasi
- Paymentgateway sudah bisa dilakukan tetapi untuk callbacknya harus online ? solusinya adalah dengan menggunakan ngrok
- copy link ngrok lalu paste kan di file constant pada variable CallbacURL
- ikuti sesuai dengan postman
