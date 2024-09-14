## Yakuniy Imtihon loyihasi: Personal Finance Management System ##

## 1. Umumiy korinishi

Mikroservislar arxitekturasidan foydalangan holda Shaxsiy moliyani boshqarish tizimini yarating. Ushbu tizim foydalanuvchilarga o'z daromadlari va xarajatlarini kuzatish, byudjetlarni boshqarish va moliyaviy hisobotlarni yaratish imkonini beradi. Loyiha mikroservislar, konteynerlashtirish, asinxron aloqa va real dunyo dizayn tamoyillari bilan amaliy tajribani taqdim etadi.

## 2. Arxitektura komponentlari

1. Income & Expenses Service - Foydalanuvchilarning daromadlari va xarajatlarini kuzatib boradi.
2. Budget Service - Foydalanuvchilarga byudjetlarni yaratish va nazorat qilish imkonini beradi.
3. User Service - Foydalanuvchilarni ro'yxatdan o'tkazish, autentifikatsiya qilish va profilni boshqarish bilan shug'ullanadi.
4. Notification Service - Byudjetlar oshib ketganda yoki xarajatlar chegaralari bajarilganda bildirishnomalarni yuboradi.
5. `API Gateway` - So‘rovlarni tegishli microservice larga yo‘naltiradi.
6. Reporting Service - Foydalanuvchining moliyaviy ma'lumotlari asosida hisobotlarni yaratadi.

3. Microservice API Endpointlari
3.1 Income & Expenses Service
1. Endpoint: `POST /api/v1/transactions/income`
* Tavsif: Yangi daromad operatsiyasini ro'yxatdan o'tkazing.
* So’rov:
{
  "amount": 5000,
  "currency": "USD",
  "category": "Daromad",
  "date": "2024-09-13"
}
​
Javob: 
{
  "message": "Income logged successfully",
  "transactionId": "tx123"
}
​
2. Endpoint: `POST /api/v1/transactions/expense`
* Tavsif:  Yangi xarajat tranzaksiyasini qo‘shish.
* So’rov:
{
  "amount": 100,
  "currency": "USD",
  "category": "Oziq-ovqat",
  "date": "2024-09-13"
}
​
Javob: 
{
  "message": "Xarajat muvaffaqiyatli kiritildi",
  "transactionId": "tx124"
}
​
3. Endpoint: `GET /api/v1/transactions`
* Tavsif: Foydalanuvchi uchun barcha daromad va xarajatlarni olish.
* Javob:
[
  {
    "transactionId": "tx123",
    "type": "Daromad",
    "amount": 5000,
    "currency": "USD",
    "category": "Oylik",
    "date": "2024-09-13"
  },
  {
    "transactionId": "tx124",
    "type": "Xarajat",
    "amount": 100,
    "currency": "USD",
    "category": "Oziq-ovqat",
    "date": "2024-09-13"
  }
]

3.2 Budget Service 
1. Endpoint: `POST /api/v1/budgets`
* Tavsif: Ma'lum bir kategoriya uchun yangi byudjet yaratish.
* So'rov:
{
  "category": "Oziq-ovqat",
  "amount": 300,
  "currency": "USD"
}
​
Javob:
{
  "message": "Byudjet muvaffaqiyatli yaratildi",
  "budgetId": "bud123"
}
​
2. Endpoint: `GET /api/v1/budgets`
* Tavsif: Foydalanuvchi uchun barcha byudjetlarni olish.
* Javob:
[
  {
    "budgetId": "bud123",
    "category": "Oziq-ovqat",
    "amount": 300,
    "spent": 100,
    "currency": "USD"
  }
]
​
3. Endpoint: `PUT /api/v1/budgets/{budgetId}`
Tavsif: Mavjud byudjetni yangilash.
So'rov:
{
  "amount": 400
}
​
Javob:
{
  "message": "Byudjet muvaffaqiyatli yangilandi"
}

3.3 User Service
1. Foydalanuvchi Ro'yxatdan O'tishi
* Endpoint: `POST /api/v1/users` 
* Tavsif: Yangi foydalanuvchini ro‘yxatdan o‘tkazish.
* So'rov:
{
   "username": "string",
   "password": "string",
   "email": "string"
}
​
Javob:
 {
    "userID": "string",
    "username": "string",
    "email": "string"
 }
​
2. Foydalanuvchi Kirishi
* Endpoint: `POST /api/v1/users/login`
* Tavsif: Foydalanuvchi kirishi va token olishi.
* So'rov:
{
   "username": "string",
   "password": "string"
}
​
Javob:
 {
    "token": "string",
    "expiresIn": "number"// Tokenning amal qilish muddati (sekundlarda)
 }
​
3. Foydalanuvchi profili ma'lumotlarini olish:
* Endpoint: `GET /api/v1/users/profile`
* Javob:
{
  "userId": "userId",
  "username": "username",
  "email": "email@example.com"
}

​
3.4 Reporting Service
1. Endpoint: `GET /api/v1/reports/income-expense` 
* Tavsif: Daromad va xarajatlar bo‘yicha umumiy hisobotni olish.
* Javob:
{
  "totalIncome": 5000,
  "totalExpenses": 100,
  "netSavings": 4900
}
​
2. Endpoint: `GET /api/v1/reports/spending-by-category`
* Tavsif: Kategoriyalar bo‘yicha xarajat hisobotini olish.
* Javob:
[
  {
    "category": "Oziq-ovqat",
    "totalSpent": 100
  },
  {
    "category": "O‘yin-kulgi",
    "totalSpent": 50
  }
]

## 4. Tizim oqimi

1. **Foydalanuvchini ro'yxatdan o'tkazish va autentifikatsiya qilish**:
    * Foydalanuvchilar `API Gateway` orqali Foydalanuvchi xizmati orqali ro'yxatdan o'tadi va autentifikatsiya qiladi.
    * Foydalanuvchi xizmati kelgusida autentifikatsiya qilingan so'rovlar uchun `JWT` tokenini chiqaradi.
    * Service lar:  `API Gateway` → `User Service` → `PostgreSQL`

2. **Daromad va xarajatlarni hisobga olish**:
    * Foydalanuvchilar `API Gateway` orqali daromad va xarajatlarni qayd qiladilar va daromadlar va xarajatlar xizmati bu operatsiyalarni qayd qiladi.
    * Agar foydalanuvchi byudjet chegaralariga yaqinlashsa, tizim bildirishnomalarni ishga tushiradi.
    * Service lar: `API Gateway` → `Income & Expenses Service` → `Notification Service` → `PostgreSQL`

3. **Byudjetlarni boshqarish:**
    * Foydalanuvchilar  `API Gateway`orqali byudjetlar yaratadilar yoki boshqaradilar va Budget Service belgilangan byudjetlarga nisbatan sarf-xarajatlarni kuzatib boradi.
    * Cheklovlar o’shib ketganda, notification service foydalanuvchini ogohlantiradi.
    * Service lar `API Gateway`→ `Budget Service` → `Redis` → `Notification Service`

4. **Hisobotlarni ko‘rish**:
    * Foydalanuvchilar `API Gateway` orqali hisobotlarni so'rashadi va Reporting Service daromadlar va xarajatlar va byudjet service ma'lumotlar bazalaridan hisobotlarni yaratadi.
    * Services: `API Gateway` → `Reporting Service` → `PostgreSQL/Redis`

5. Real vaqtda bildirishnomalar:
    * Tizim `WebSocket` yordamida byudjet chegaralariga yaqinlashganda real vaqtda ogohlantirishlarni yuboradi.
    * Services: `API Gateway`→ `Income & Expenses Service` → `Notification Service` → `RabbitMQ/Kafka`

## 5. Texnologiyalar va Talablar

* **Microservicelar Aloqasi**: Microservicelar orasida aloqani o‘rnatish uchun `gRPC` foydalaniladi.
* **Message Brokeri**: `Kafka` yoki `RabbitMQ` asinxron aloqa va xabarlar uchun ishlatiladi.
* **Real Vaqtda Xabarlar**:  Real vaqt xabarlari uchun `WebSockets` ishlatiladi.
* **API Gateway**:  So‘rovlarni tegishli microservicelarga yo‘naltiradi.
* **Swagger**: API hujjatlari va testlari uchun foydalaniladi.
* **Email Xabarlari**: Harajatlar mal’um bir chegaradan o’shib ketdanga email orqali yuborish.
* `HTTPS`: Xavfsiz aloqa uchun `HTTPS` ishlatiladi.
* `Rate Limiting`: Suiste'molni oldini olish uchun `rate limiting` amalga oshiriladi.
* `Graceful Shutdown`: Servicelarni to‘g‘ri to‘xtatilishini ta'minlash.
* **Konfiguratsiya Boshqaruvi**: Konfiguratsiyalarni muhit o‘zgaruvchilari yoki konfiguratsiya boshqaruv tizimi yordamida boshqarish.

## 6. Baholash uchun ballar

- 

## 7.  Big picture