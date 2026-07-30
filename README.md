# Golang Mock API Server (Lightweight & Docker Ready)

Bu proje, mobil/web veya arka plan uygulamanızı test edebilmeniz için geliştirilmiş **çok düşük kaynak tüketen** ve sıfır harici kütüphane bağımlılığı olan bir Golang Mock API servisidir.

## 🚀 Özellikler

- **Düşük Kaynak Tüketimi**: Standart Go kütüphanesi (`net/http`) ile yazılmıştır, dış bağımlılık içermez. Multi-stage Docker build sayesinde nihai image boyutu **~10 MB**, çalışma anındaki RAM tüketimi ise **~4 MB** civarındadır.
- **Kota & Limit Yok**: İster günde 1 istek, ister 1 milyon istek atın; herhangi bir günlük istek kotası veya rate limit bulunmaz.
- **CORS Desteği**: Tarayıcı / webview üzerinden gelen preflight (`OPTIONS`) isteklerini destekler (`Access-Control-Allow-Origin: *`).
- **Doğrudan JSON Array**: `/api/v1/children` ve `/api/v1/transactions` endpoint'leri hiçbir wrapper nesneye sarılmadan doğrudan JSON dizisi (`[...]`) döndürür.
- **Content-Type**: Tüm başarılı ve hatalı yanıtlarda `Content-Type: application/json` başlığı otomatik set edilir.

---

## 📌 Desteklenen Endpoint'ler ve Senaryolar

Her bir endpoint aşağıdaki query parametrelerini destekler:
- `?scenario=success` (veya query gönderilmediğinde varsayılan)
- `?scenario=empty`
- `?scenario=error` (HTTP 500 kodu ile JSON hata mesajı döner)

### 1. Wallet Endpoint
```http
GET /api/v1/wallet
GET /api/v1/wallet?scenario=success
GET /api/v1/wallet?scenario=empty
GET /api/v1/wallet?scenario=error
```

### 2. Children Endpoint (JSON Array)
```http
GET /api/v1/children
GET /api/v1/children?scenario=success
GET /api/v1/children?scenario=empty
GET /api/v1/children?scenario=error
```

### 3. Transactions Endpoint (JSON Array)
```http
GET /api/v1/transactions
GET /api/v1/transactions?scenario=success
GET /api/v1/transactions?scenario=empty
GET /api/v1/transactions?scenario=error
```

---

## 🌐 Caddy ile Kullanım (`saas_default` Ağı)

Projedeki `docker-compose.yml` dosyası, sunucunuzdaki Caddy ile aynı ağa bağlanacak şekilde **`saas_default`** dış ağına (`external: true`) ayarlanmıştır.

> **Önemli:** Caddy zaten `saas_default` ağı içinden doğrudan `mock-api:8080` ile iletişim kurabildiği için, sunucunuzdaki diğer servislerle herhangi bir port çakışması (`port is already allocated`) olmaması adına dışarıya (host / `0.0.0.0`) hiçbir port açılmamıştır. Tüm trafik güvenli bir şekilde Caddy üzerinden HTTPS ile yönlendirilir.

### Caddyfile Ayarı
Sunucunuzdaki `Caddyfile` dosyanıza aşağıdaki bloğu eklemeniz yeterlidir:

```caddyfile
# Mock API Servisi
mock.basitemlak.com {
    reverse_proxy mock-api:8080
}
```

Caddyfile dosyasını kaydettikten sonra Caddy'yi yenileyin:
```bash
docker compose exec caddy caddy reload
```

---

## 🐙 GitHub'a Yükleme ve Sunucuda Deploy Etme

### 1. Adım: Projeyi GitHub'a Yükleme (Kendi Bilgisayarınızdan)
Terminalinizde projenin bulunduğu `/Users/yusufkalyoncu/Desktop/mock-api` klasöründe aşağıdaki komutları çalıştırın:

```bash
# Git reposu başlatın
git init

# Dosyaları ekleyin (.gitignore gereksiz dosyaları otomatik dışarıda bırakacaktır)
git add .

# İlk commit'inizi alın
git commit -m "Initial commit: Lightweight Golang Mock API"

# Ana dal adını main yapın
git branch -M main

# GitHub'da oluşturduğunuz boş reponun adresini bağlayın
git remote add origin https://github.com/<kullanici-adiniz>/<repo-adi>.git

# GitHub'a yükleyin
git push -u origin main
```

---

### 2. Adım: Sunucuda Çekme ve Deploy Etme (Sunucunuzda)
Sunucunuza SSH ile bağlandıktan sonra:

```bash
# 1. GitHub'dan projeyi çekin
git clone https://github.com/<kullanici-adiniz>/<repo-adi>.git

# 2. Proje klasörüne girin
cd <repo-adi>

# 3. Docker Compose ile derleyip arka planda çalıştırın
docker compose up -d --build
```

Sunucunuzda mock API artık **`https://mock.basitemlak.com`** adresinde çalışıyor! 🎉
*(Örn: `https://mock.basitemlak.com/api/v1/wallet`)*

#### Güncelleme Yaparsanız Sunucuda Yenilemek İçin:
Eğer kodda değişiklik yapıp GitHub'a tekrar push ederseniz, sunucuda sadece şu iki komutu çalıştırmanız yeterlidir:
```bash
git pull
docker compose up -d --build
```

---

### 3. Adım: İşiniz Bitince Sunucudan Tamamen Temizleme (Uçurma)

İşiniz bitip sunucuyu eski tertemiz haline getirmek istediğinizde projenin olduğu klasörün içindeyken:

```bash
# 1. Container'ı durdurup silin ve oluşturulan 10MB'lık docker image'ını da kaldırın
docker compose down --rmi all

# 2. Bir üst klasöre çıkıp proje klasörünü tamamen silin
cd .. && rm -rf <repo-adi>
```
Bu işlemlerden sonra sunucunuzda projenin ne container'ı, ne image'ı, ne de kaynak kodu kalır—tamamen uçurulmuş olur!

---

## 🧪 Test Komutları (cURL)

Aşağıdaki komutlarla canlıda veya yerelde API yanıtlarını test edebilirsiniz:

```bash
# Wallet (Success)
curl -i https://mock.basitemlak.com/api/v1/wallet?scenario=success

# Children (Success - Doğrudan JSON Array)
curl -i https://mock.basitemlak.com/api/v1/children?scenario=success

# Transactions (Success - Doğrudan JSON Array)
curl -i https://mock.basitemlak.com/api/v1/transactions?scenario=success

# Empty Senaryoları
curl -i https://mock.basitemlak.com/api/v1/wallet?scenario=empty
curl -i https://mock.basitemlak.com/api/v1/children?scenario=empty
curl -i https://mock.basitemlak.com/api/v1/transactions?scenario=empty

# Error Senaryosu (HTTP 500 döner)
curl -i https://mock.basitemlak.com/api/v1/wallet?scenario=error
```
