# Mock API Dökümantasyonu (API Reference & Response Models)

Bu döküman, **Mock API** servisinin sunduğu tüm endpoint'leri, desteklenen sorgu parametrelerini, HTTP durum kodlarını, JSON yanıt modellerini ve öncelikli olarak web/mobil önyüzlerde kullanılmak üzere hazırlanmış **TypeScript/JavaScript veri modellerini (interface)** içerir.

---

## 🌐 Base URL (Temel Adres)

| Ortam | Base URL |
| :--- | :--- |
| **Canlı (Caddy / Sunucu)** | `https://mock.basitemlak.com` |
| **Yerel (Localhost)** | `http://localhost:8080` |

---

## 🎯 Ortak Özellikler ve Senaryo Yönetimi (`?scenario`)

Tüm GET endpoint'leri davranışlarını dinamik olarak test edebilmeniz için `scenario` query parametresini destekler:

| Parametre | Tipi | Varsayılan | Açıklama |
| :--- | :--- | :--- | :--- |
| `scenario` | `string` | `success` | Yanıt davranışını belirler. Geçerli değerler: `success`, `empty`, `error` |

- **`?scenario=success`**: Standart, örnek verilerle dolu başarılı yanıtı döner (HTTP 200). Parametre gönderilmediğinde de varsayılan olarak bu çalışır.
- **`?scenario=empty`**: Liste endpoint'lerinde boş dizi (`[]`), cüzdan endpoint'inde `0.00` bakiye döner (HTTP 200).
- **`?scenario=error`**: Sunucu hatası simüle etmek için HTTP 500 kodu ile standart hata nesnesi döner.

---

## 📦 TypeScript / JavaScript Veri Modelleri (Models & Interfaces)

Aşağıdaki interface'leri Next.js, React, React Native veya herhangi bir TypeScript/JavaScript projenizde doğrudan kopyalayıp kullanabilirsiniz:

```typescript
// 1. Cüzdan Modeli (Wallet)
export interface WalletResponse {
  id: string;          // Cüzdan benzersiz ID'si (Örn: "wal_29a1f3c8")
  currency: string;    // Para birimi (Örn: "TRY")
  balance: number;     // Güncel bakiye (Örn: 1250.75)
  created_at: string;  // ISO-8601 formatında tarih saati
}

// 2. Çocuk Profili Modeli (Child)
export interface ChildResponse {
  id: string;             // Çocuk benzersiz ID'si (Örn: "child_001")
  first_name: string;     // Adı
  last_name: string;      // Soyadı
  avatar_url: string | null; // Avatar resim bağlantısı (null olabilir)
  age: number;            // Yaş (Örn: 12)
  grade: string;          // Sınıf bilgisi (Örn: "7th Grade")
  wallet_balance: number; // Çocuğa tanımlı cüzdan bakiyesi
  currency: string;       // Para birimi (Örn: "TRY")
  school_name: string;    // Okul adı
}

// 3. İşlem Geçmişi Modeli (Transaction)
export type TransactionType = "income" | "expense";
export type TransactionCategory = "cafeteria" | "topup" | "trip" | "campus_store" | "event" | "cashback";
export type TransactionStatus = "completed" | "pending" | "failed";

export interface TransactionResponse {
  id: string;                // İşlem benzersiz ID'si (Örn: "txn_a001")
  type: TransactionType;     // İşlem türü: "income" (gelir) | "expense" (gider)
  category: TransactionCategory; // İşlem kategorisi
  description: string;       // İşlem açıklaması
  amount: number;            // Miktar (Giderlerde negatif, gelirlerde pozitif)
  currency: string;          // Para birimi (Örn: "TRY")
  date: string;              // ISO-8601 formatında işlem saati
  child_id: string | null;   // İşlemin ilişkili olduğu çocuk ID'si (Genel yüklemeler için null)
  status: TransactionStatus; // İşlem durumu
}

// 4. Standart Hata Yanıtı Modeli (Error Response - HTTP 500)
export interface ErrorResponse {
  code: string;    // Hata kodu (Örn: "INTERNAL_SERVER_ERROR")
  message: string; // Açıklayıcı hata mesajı
}
```

---

## 🛠️ Endpoint Referansı

### 1. Cüzdan Bilgisi Getirme

Kullanıcının ana cüzdan bakiyesini, para birimini ve hesap oluşturulma tarihini getirir.

- **URL:** `/api/v1/wallet`
- **Method:** `GET`
- **Content-Type:** `application/json`

#### Yanıtlar:

##### ✅ Başarılı Yanıt — `HTTP 200 OK` (`?scenario=success` veya varsayılan)
```json
{
  "id": "wal_29a1f3c8",
  "currency": "TRY",
  "balance": 1250.75,
  "created_at": "2025-09-15T10:30:00Z"
}
```

##### 📭 Boş Yanıt — `HTTP 200 OK` (`?scenario=empty`)
```json
{
  "id": "wal_new_user_01",
  "currency": "TRY",
  "balance": 0.00,
  "created_at": "2026-03-30T08:00:00Z"
}
```

##### ❌ Hata Yanıtı — `HTTP 500 Internal Server Error` (`?scenario=error`)
```json
{
  "code": "INTERNAL_SERVER_ERROR",
  "message": "The requested information could not be loaded."
}
```

---

### 2. Çocuk Listesi Getirme

Hesaba bağlı çocukların profillerini, okullarını ve cüzdan bakiyelerini listeler. **Herhangi bir wrapper nesneye sarılmadan doğrudan JSON Array (`ChildResponse[]`) döndürür.**

- **URL:** `/api/v1/children`
- **Method:** `GET`
- **Content-Type:** `application/json`

#### Yanıtlar:

##### ✅ Başarılı Yanıt — `HTTP 200 OK` (`?scenario=success` veya varsayılan)
```json
[
  {
    "id": "child_001",
    "first_name": "Elif",
    "last_name": "Yılmaz",
    "avatar_url": null,
    "age": 12,
    "grade": "7th Grade",
    "wallet_balance": 340.00,
    "currency": "TRY",
    "school_name": "Bahçeşehir Koleji"
  },
  {
    "id": "child_002",
    "first_name": "Can",
    "last_name": "Yılmaz",
    "avatar_url": null,
    "age": 9,
    "grade": "4th Grade",
    "wallet_balance": 125.50,
    "currency": "TRY",
    "school_name": "Bahçeşehir Koleji"
  },
  {
    "id": "child_003",
    "first_name": "Zeynep",
    "last_name": "Yılmaz",
    "avatar_url": null,
    "age": 15,
    "grade": "10th Grade",
    "wallet_balance": 580.00,
    "currency": "TRY",
    "school_name": "Doğa Koleji"
  }
]
```

##### 📭 Boş Yanıt — `HTTP 200 OK` (`?scenario=empty`)
```json
[]
```

##### ❌ Hata Yanıtı — `HTTP 500 Internal Server Error` (`?scenario=error`)
```json
{
  "code": "INTERNAL_SERVER_ERROR",
  "message": "The requested information could not be loaded."
}
```

---

### 3. İşlem Geçmişi Getirme

Hesaptaki tüm harcama, yükleme ve iade işlemlerini tarihe göre listeler. **Herhangi bir wrapper nesneye sarılmadan doğrudan JSON Array (`TransactionResponse[]`) döndürür.**

- **URL:** `/api/v1/transactions`
- **Method:** `GET`
- **Content-Type:** `application/json`

#### Yanıtlar:

##### ✅ Başarılı Yanıt — `HTTP 200 OK` (`?scenario=success` veya varsayılan)
```json
[
  {
    "id": "txn_a001",
    "type": "expense",
    "category": "cafeteria",
    "description": "School Cafeteria — Elif",
    "amount": -45.00,
    "currency": "TRY",
    "date": "2026-03-28T12:35:00Z",
    "child_id": "child_001",
    "status": "completed"
  },
  {
    "id": "txn_a002",
    "type": "income",
    "category": "topup",
    "description": "Wallet Top-Up",
    "amount": 500.00,
    "currency": "TRY",
    "date": "2026-03-27T09:15:00Z",
    "child_id": null,
    "status": "completed"
  },
  {
    "id": "txn_a003",
    "type": "expense",
    "category": "trip",
    "description": "Cappadocia School Trip — Can",
    "amount": -1200.00,
    "currency": "TRY",
    "date": "2026-03-25T16:00:00Z",
    "child_id": "child_002",
    "status": "completed"
  },
  {
    "id": "txn_a004",
    "type": "expense",
    "category": "cafeteria",
    "description": "School Cafeteria — Zeynep",
    "amount": -32.50,
    "currency": "TRY",
    "date": "2026-03-25T12:20:00Z",
    "child_id": "child_003",
    "status": "completed"
  },
  {
    "id": "txn_a005",
    "type": "expense",
    "category": "campus_store",
    "description": "Campus Bookstore — Elif",
    "amount": -89.90,
    "currency": "TRY",
    "date": "2026-03-24T14:10:00Z",
    "child_id": "child_001",
    "status": "completed"
  },
  {
    "id": "txn_a006",
    "type": "income",
    "category": "topup",
    "description": "Wallet Top-Up",
    "amount": 1000.00,
    "currency": "TRY",
    "date": "2026-03-22T08:45:00Z",
    "child_id": null,
    "status": "completed"
  },
  {
    "id": "txn_a007",
    "type": "expense",
    "category": "event",
    "description": "Science Fair Registration — Can",
    "amount": -150.00,
    "currency": "TRY",
    "date": "2026-03-20T10:30:00Z",
    "child_id": "child_002",
    "status": "completed"
  },
  {
    "id": "txn_a008",
    "type": "expense",
    "category": "cafeteria",
    "description": "School Cafeteria — Elif",
    "amount": -38.00,
    "currency": "TRY",
    "date": "2026-03-19T12:45:00Z",
    "child_id": "child_001",
    "status": "completed"
  },
  {
    "id": "txn_a009",
    "type": "income",
    "category": "cashback",
    "description": "Cashback Reward — March",
    "amount": 25.00,
    "currency": "TRY",
    "date": "2026-03-18T00:00:00Z",
    "child_id": null,
    "status": "completed"
  },
  {
    "id": "txn_a010",
    "type": "expense",
    "category": "trip",
    "description": "Visa Fee — Italy Trip — Zeynep",
    "amount": -320.00,
    "currency": "TRY",
    "date": "2026-03-15T11:00:00Z",
    "child_id": "child_003",
    "status": "completed"
  }
]
```

##### 📭 Boş Yanıt — `HTTP 200 OK` (`?scenario=empty`)
```json
[]
```

##### ❌ Hata Yanıtı — `HTTP 500 Internal Server Error` (`?scenario=error`)
```json
{
  "code": "INTERNAL_SERVER_ERROR",
  "message": "The requested information could not be loaded."
}
```

---

## 🏥 Health Check (Durum Kontrolü)

Container'ın veya servisin hayatta olup olmadığını kontrol etmek için kullanılır.

- **URL:** `/health`
- **Method:** `GET`
- **Yanıt (`HTTP 200 OK`):**
```json
{
  "status": "ok"
}
```
