# System Design Test

## Soal
Anda diminta untuk merancang sebuah sistem E-Ticketing Transportasi Publik yang
beroperasi 24 jam, dimana tarif ditentukan berdasarkan titik terminal ketika checkin
dan titik terminal ketika checkout. Untuk titik terminal ada 5 dan mempunyai
gate/gerbang validasi yang terhubung pada server (bisa 1 atau lebih). Pembayaran
dilakukan menggunakan kartu prepaid (seperti pada tol) yang mampu berjalan saat
offline.
1. Gambarkan desain rancangan anda
2. Ceritakan rancangan anda dengan jelas saat ada jaringan internet
3. Ceritakan solusi anda dengan jelas (apabila memungkinkan) saat tidak ada
jaringan internet


# Solutions
## Problem Statement & Requirements

### Business Requirements

-   **Operasional 24/7**: Sistem harus dapat beroperasi tanpa henti
-   **Distance-based Fare**: Tarif dihitung berdasarkan terminal asal dan tujuan
-   **5 Terminal**: Dengan multiple gate per terminal untuk throughput tinggi
-   **Prepaid Card System**: Menggunakan kartu berisi saldo seperti sistem tol
-   **Offline Capability**: Tetap dapat beroperasi saat koneksi internet terputus

### Technical Requirements

-   **High Availability**: Minimal 99.9% uptime
-   **Low Latency**: Response time < 500ms untuk tap kartu
-   **Scalability**: Support pertumbuhan penumpang dan terminal
-   **Data Consistency**: Saldo kartu harus akurat dan konsisten
-   **Security**: Proteksi terhadap fraud dan double charging

## 1. Desain Rancangan Sistem

![enter image description here](https://mermaid.ink/img/pako:eNqdWN1u2zYUfhVCQHsVJ41jJ40uCkSS0x_kr42LAUt2QUuMTVgWPVJK6sUFdrGLYb9d24uhwDAUGLBhu-kTbO_SJ9gj7JCULOonmp1cBNLh9x2SRzzfOfS15bOAWLZ1EbIrf4R5jPreeYTgTySDIcfTEXouCD_AM8LRmXxEd5FHLqlPkDJ-ptHyT46e_fvrT7-pJ3RF4xE62neRi3mwccgGNCQGuk_4hEY43APKuz8Xr2gPPcRxHdApAZ2bgG4J6N4E9EpA7yZgrwTsmUASBedRKWaPoyEnQqRheyxiytAp4Spsh0SMKrFTkJQFk333SlsyR2q-KzwzGKk76Q0Ib79ENbOc0oD4mIuGlTqJoFG-1Ow18yMqK-1Tf0xiGg1TBEz-w1-5NSMajH3MSQ5-_UEZanBHLKYX1MewjSjHv_m9MFDDe0RFzPjMoHyb2dATlvCIzGpYe_AhZzH1hcH7JrfWME5nke_hGOeEt18pI5LWEqEu1hKWxlkzYIl4WE2kZySg4uzju_f6CRLIH5kLOWEihmNx-vQA1vDqZ-MdnXA6wbBvzzEX_vSAxkQm2uu_zUTT9grQKQGdm4BuCejeBPRKQO8mYK8E7FWBcLAFBM3hbKzk5s0fmQlpm4F1sD9OpmmYpet_UlMW-oaPdTwAEbvEoFo0nqVfrWADITxkEQVHcOyNSY-nJOqTkExIzGdyge-LJuSyMCQ-8MxPyhkMjkgis__N18a7AXqCyVDt-e336bMxeMDGVHJ_UU_GwEOOL3CE9flOX-AoitGAgTA3RMDlLHrCBgKdncL5C5KQBEi-VwIsgWcff_yQBTclVl3fuaNLQ8zyT7wPlUePqqFW68G8j6eLwjHPC8VSKGcplLsUylsK1TN2t9gV7FDrcb69xT6Ur0f9_gmk9-cJEfG8IP9FtLMS2l0J7a2E7v0vehEGvfUjxuNR65QlcSmbNM7kotY6-DYJfTiooPnzYu5U5ogXFU_UeW09MOukRhgGBSgXNPNzlquaVnoQWUh6UADj45acqFhtrmvpRvuUy2gpNW_At9dBt6E6huEAEmlu6HoDaWsdHV9chFC1QY4CMs_UfnWKszrFXZ3irU4xM0w2QqCgUcxZCAEVUxYJ0uBSHp8NN2SCKGpFTlYjObchubchebchmYHaE7Iz6WMxFuiS4lKRbPB-kgxCCrnRuyRRLObFgqt5-n9hIBVIMUabNtSBEEfg7ngKfKP9a-S1bciyKEC9CabhvK4bbKRv2Ytm74AN56WusJHasfO-b15pDBupXTtt_iDY83KHaHyNRTsNgiURA5ydWiM8WmBZGGy4OPSTUB71NJRVMaiJj-LD5gtjokothkaxnhGf8SCLYJVSDooiySaKmKEzWfnWyx2ylHSu15YqcjFs6YJwkHWz1eXUMZ5PA6UNUnELWlsHVotKtaYkmcvCndXg7mpwbzW4mfmZhMpvekl4Wm7TPeoUhzyTVQ2KbCSwnx6Tyuk1rgS3o7m3o3m3o_VWoS2i1cMibn0C3Uxdj1IVSNmnHEIrIq-Id-UkkNSVJqWS10uzarN6aXY5sZcmVtN7aWrlhC7FzA9r4Up1QqdEnl09XLw2qaMvC0TqfG5ckhrx2RL0lakRCuIJQHmBMgtdPo9qGNM7lB7TTqt26aNoXWw5vSWBV3_RZOcXKb2Y9MKV2qsCaFz5cyXOwIVrr7VmDTkNLDvmCVmzJtAqYPlqXUtP5xbsa0LOLRseA8zH59Z59BI4Uxx9ytgko3GWDEeWfYFDAW-JklqPYrgo5hAIIOEuS6LYsrs728qHZV9bLyy7tXmv21nvdDr325325v2d7po1s-z2-r3O7m63vdPpdrZ2ttvtl2vWF2pSGNltb2_tdLu7O53dre0t8AaCDjs61L9Xqp8tX_4HlnkWfw?type=png)

[Diagram Link For Better View](https://mermaid.live/view#pako:eNqdWMtu20YU_ZUBgWZlKZIfsa1FAJOU84AdO7GCArW7GJFjaSCKo86QTtUoQBddFH2mSRZFgKIIUKBFu8kXtP-SL-gn9M6D4vBhVrIXBnnnnDszl3PPvaPnTsBC4vScy4g9C8aYJ2jgX8QI_kQ6HHE8G6OngvAjPCccnctHdAv55IoGBCnjpxot_-To-b-__vSbekLPaDJGjw495GEe3j5mQxoRCz0gfEpjHB0A5e2fy1d0gO7hpA7oloDudUCvBPSuA_oloH8dsF8C9m0gicOLuBSzB_GIEyFM2B6IhDJ0RrgK2zER40rsFMSwYLLvXmpL5kjN9wzPLYZxJ70B4c2XqGaWMxqSAHPRsFI3FTTOl5q9Zn5EZaUDGkxIQuORQcDkP_yVWzOixTjEnOTgV--VoQb3iCX0kgYYthHn-Ne_FwZqePepSBifW5RvMxt6yFIek3kN6wA-5DyhgbB43-TWGsbZPA58nOCc8OYrZUTSWiLUxVrCTJw1A5aIR9VEekJCKs4_vH2nnyCBgrG9kFMmEjgWZ4-PYA0vf7be0SmnUwz79l174Y-PaEJkor362040ba8A3RLQvQ7olYDedUC_BPSvA_ZLwH4VCAdbQNBcziZKbl7_kZmQtllYFweTdGbCLF3_Y0xZ6Bs-1skQROwKg2rRZG6-WsEGQnjMYgqO4Nhbk57MSDwgEZmShM_lAt8VTchjUUQC4NmflDMYHJNUZv_rr613C_QQk5Ha85vvzbM1eMQmVHJ_UU_WwD2OL3GM9fk2L3AUxXjIQJgbIuBxFj9kQ4HOz-D8hWlEQiTfKwGWwPMPP77PgmuIVdcffaRLQ8LyT3wIlUePqqFW6-5igGfLwrHIC8VKKHcllLcSyl8J1bd2t9wV7FDrcb695T6Ur_uDwSmk92cpEcmiIP9FtLsW2lsL7a-F7v8vehkGvfVHjCfj1hlLk1I2aZzNRa02-LYJAziooPmLYu5U5kiWFU_UeW3dteukRlgGBSgXNPtzlquaVnoQWUh6UADr45acqFh121q60SHlMlpKzRvwm23QbaiOUTSERFpYut5A2mqjk8vLCKo2yFFIFpnar09x16d461P89Sl2hslGCBQ0TjiLIKBixmJBGlzK43Pbi5ggilqRk_VI7k1I3k1I_k1IdqAOhOxMBlhMBLqiuFQkG7yfpsOIQm70r0iciEWx4Gqe_l8YMAIpJqjbgzoQ4RjcncyAb7V_jbzNHmRZHKL-FNNoUdcNNtK3estm74iNFqWusJG63cv7vkWlMWyk7vRM8wfBXpQ7ROtrLNtpECyJGOLs1Frh0QLLovC2h6MgjeRRN6GsikFNfBQfNl8YE1VqMTSK9YQEjIdZBKuUclAUSTZRxA6dzcq3Xu6QpaRzvTajyMWwmQXhMOtmq8upYzydhUobpOIWtLYOrBZltKYkmavC3fXg3npwfz24nfmZhMpvekW4KbdmjzrFIc9kVYMiGwscmGNSOb3WleBmNO9mNP9mtP46tGW0-lgkrY-hm6nrUaoCKfuUY2hF5BXxlpwEkrrSpFTyemVWbVavzC4n9srEanqvTK2c0JWY-WEtXKlO6YzIs6uHi9cmdfRlgTDOF9YlqRGfLUFfmRqhIJ4AlBcou9Dl86iG0dyh9Jh2WrVLH0XrcsvmlgReg2WTnV-k9GLMhcvYqwJoXflzJc7AhWuvs-GMOA2dXsJTsuFMoVXA8tV5Lj1dOLCvKblwevAYYj65cC7iF8CZ4fgTxqYZjbN0NHZ6lzgS8JYqqfUphotiDoEAEu6xNE6c3t7WpvLh9J47nzu9VrfbbXc6nW7nzl63u7O5tb294cwBttfe3N3d7ezv3-nsb-3uvNhwvlCzdtvbd7b3N3c6nb2d_T0Y2nBA0GFHx_r3SvWz5Yv_AI5qFm8)

# E-Ticketing System Flow Description

## 📋 System Overview

E-Ticketing system untuk transportasi publik dengan 5 terminal (A, B, C, D, E) yang beroperasi 24 jam dengan kemampuan offline. Sistem menggunakan kartu prepaid dengan tarif berdasarkan terminal asal dan tujuan.

### User Data Structure
```sql
users {
    id: PRIMARY KEY
    full_name: VARCHAR
    card_id: VARCHAR UNIQUE
    balance: DECIMAL
    registered_at: TIMESTAMP
    status: ENUM('active', 'blocked')
    current_trip_terminal_id: INT NULL  -- Track active journey
}
```

### Fare Matrix and Fare data Structure
```
fares {
    id: PRIMARY KEY
    origin_terminal_id: INT REFERENCES terminals(id)
    destination_terminal_id: INT REFERENCES terminals(id)
    price: DECIMAL
    is_active: BOOLEAN
}
```

| Checkin → Checkout | A    | B    | C    | D    | E    |
|-------------------|------|------|------|------|------|
| A                 | 0    | 500  | 700  | 900  | 1200 |
| B                 | 500  | 0    | 600  | 800  | 1100 |
| C                 | 700  | 600  | 0    | 700  | 1000 |
| D                 | 900  | 800  | 700  | 0    | 900  |
| E                 | 1200 | 1100 | 1000 | 900  | 0    |

---

## 🔄 Flow 1: Check-In Process (Online)

### Step 1: User Interaction
1. **User** tap kartu NFC atau mobile dengan NFC feature ke **Gate Device Entrance Terminal A**
2. Gate Device capture card data dan buat HTTP request ke **CheckIn endpoint**

### Step 2: Request Routing & Observability
1. Gate Device request masuk ke **Istio Ingress Gateway**
2. **Istio Ingress** forward request melalui **Service Mesh** ke **Ticketing Service**
3. **Observability**: Istio Ingress kirim North-South traffic data ke **OpenTelemetry Collector**
4. OpenTelemetry distribute data ke:
   - **Prometheus** (metrics)
   - **Jaeger** (tracing)
   - **Loki** (logs)

### Step 3: User Data Retrieval (Priority Strategy)
**Ticketing Service** menggunakan data priority strategy:
1. **First Priority**: Query **Redis Cache** untuk fast lookup
   - Jika user data ditemukan → proceed to validation
2. **Second Priority**: Query **PostgreSQL Database** jika cache miss
   - Load user data dari primary database
   - Update Redis cache untuk future requests
3. **Fallback Priority**: Query **Local SQLite** jika Offline 
   - Use local user snapshot for basic validation

### Step 4: User Validation
**Ticketing Service** validate user data:
1. **User Existence Check**:
   - Jika user tidak ditemukan → Response "Not Registered" → Gate tetap **CLOSED**
   - Display message: "Kartu belum terdaftar, silakan daftar terlebih dahulu"

2. **User Status Check**:
   - Jika status = 'blocked' → Response "Unauthorized" → Gate tetap **CLOSED**
   - Display message: "Kartu diblokir, hubungi customer service"

3. **Active Journey Check**:
   - Jika `current_trip_terminal_id` IS NOT NULL → User masih dalam perjalanan
   - Response "Already Checked In" → Gate tetap **CLOSED**
   - Display message: "Anda sudah check-in, lakukan check-out terlebih dahulu"

### Step 5: Balance Validation
1. **Minimum Balance Check**:
   - Ticketing Service check apakah `user.balance >= minimum_fare` (contoh: 500)
   - Jika balance insufficient → Response "Insufficient Balance" → Gate tetap **CLOSED**
   - Display balance dan minimum required

2. **Balance Hold Strategy**:
   - Jika validation passed → Hold maximum possible fare (1200) sementara
   - Ini untuk anticipate worst-case scenario (A → E)

### Step 6: Gate Control & Response
1. **Success Response**: 
   - Ticketing Service kirim command **"OPEN GATE"** ke Terminal CheckIn
   - Gate terbuka selama 5 detik
   - Display success message dengan balance info

2. **Update Current Trip**:
   - Set `current_trip_terminal_id = A` di user record

### Step 7: Async Task Distribution
**Ticketing Service** publish event ke **Message Broker** untuk 5 async tasks:

#### Task 1: **Fare Service** - Balance Hold
1️⃣ Data Structure: `balance_holds` table
```sql
balance_holds {
    id: PRIMARY KEY
    user_id: INT REFERENCES users(id)
    terminal_id: INT REFERENCES terminals(id)
    hold_amount: DECIMAL
    status: ENUM('held','charged','released')
    created_at: TIMESTAMP
    released_at: TIMESTAMP NULL
}
```
2️⃣ Event Payload (Message Broker)
```json
{
  "task": "held",
  "user_id": "123",
  "terminal_id": "A",
  "hold_amount": 1200,
  "timestamp": "2024-01-15T08:30:00Z"
}
```
- **Fare Service** update PostgreSQL: deduct User Balance 1200 dari balance (temporary hold)
- Mark transaction sebagai "held" bukan "charged"
- released_at akan di update saat checkout

#### Task 2: **Notification Service** - Check-in Alert
1️⃣ Data Structure: `notifications` table
```sql
notifications {
    id: PRIMARY KEY
    user_id: INT REFERENCES users(id)
    type: ENUM('checkin','checkout','fare_update')
    status ENUM('pending','sent','failed')
    created_at: TIMESTAMP
    sent_at: TIMESTAMP NULL
    message TEXT  -- isi notifikasi, misal "Anda check-in di Terminal A..."
}
```
2️⃣ Event Payload (Message Broker)
```json
{
  "task": "send_checkin_notification",
  "user_id": "123",
  "terminal_id": "A",
  "held_amount": 1200
}
```
- Send email: sesuai message Contoh  "Anda check-in di Terminal A pada 08:30. Saldo ditahan sementara Rp 1.200. Saldo akan disesuaikan saat check-out berdasarkan tujuan akhir."

#### Task 3: **History Service** - Journey Record
1️⃣ Data Structure: `journey_history` table
```sql
journey_history {
    id: PRIMARY KEY
    user_id: INT REFERENCES users(id)
    checkin_terminal_id: INT REFERENCES terminals(id)
    checkin_time: TIMESTAMP
    checkout_terminal_id: INT NULL REFERENCES terminals(id)
    checkout_time: TIMESTAMP NULL
    fare_charged: DECIMAL NULL
    status: ENUM('active','completed','cancelled')  -- active = checkin ongoing
}

```
2️⃣ Event Payload (Message Broker)
```json
{
  "task": "record_checkin",
  "user_id": "123",
  "terminal_id": "A",
  "timestamp": "2024-01-15T08:30:00Z",
  "status": "checkin"
}
```
- Insert ke `journey_history` table:
  - `user_id`: 123
  - `checkin_terminal`: A  
  - `checkin_time`: 2024-01-15T08:30:00Z
  - `checkout_terminal`: NULL
  - `checkout_time`: NULL
  - `fare_charged`: NULL
  - `status`: 'active'
  
- Data ini bisa digunakan untuk Sebagai Journey History dan bisa digunakan untuk Refrensi Saat User Validation 	atau saat Checkout (Refrensi ini bersifat Optional)


#### Task 4: **Analytics Service** - Usage Data
1️⃣ **Data Structure: `analytics_events` table**
```sql
analytics_events {
    id: PRIMARY KEY
    event_type: ENUM('checkin','checkout','fare_update')
    terminal_id: INT REFERENCES terminals(id)
    user_id: INT NULL REFERENCES users(id)
    user_demographics JSONB NULL
    event_timestamp: TIMESTAMP
    created_at: TIMESTAMP DEFAULT now()
}
```
```json
{
  "task": "record_analytics",
  "event_type": "checkin",
  "terminal_id": "A",
  "timestamp": "2024-01-15T08:30:00Z",
  "user_demographics": {
      "age_group": "25-34",
      "gender": "male",
      "membership": "gold"
  }
}
```
- **Analytics Service** menerima event dari Message Broker
- Insert ke tabel `analytics_events` untuk digunakan oleh **BI / Dashboard / Forecasting**
- Store untuk business intelligence, dan bisa digunakan untuk:
  - Peak hour analysis
  - Popular terminal analysis  
  - Revenue forecasting

#### Task 5: **Sync Data Service** - Data Consistency
1️⃣ **Data Structure: `offline_tasks` table for SQLite Offline mode**  
```sql
offline_tasks {
    id: PRIMARY KEY,
    service_name: VARCHAR,
    action: VARCHAR,
    payload: JSONB,
    status: ENUM('pending','processed','failed'),
    created_at: TIMESTAMP,
    processed_at: TIMESTAMP NULL
}
```

-   **users** table (PostgreSQL)
-   **Redis Cache** (fast lookup)
-   **SQLite snapshots** di tiap terminal (offline mode)

**Redis Payload Example** (fast lookup cache)
```json
{
  "task": "sync_user_data",
  "user_id": "123",
  "updated_fields": ["current_trip_terminal_id", "balance"],
  "user_snapshot": {
    "user_id": "123",
    "full_name": "Sony Nurdianto",
    "card_id": "XYZ987",
    "balance": 5000,
    "status": "active",
    "current_trip_terminal_id": 2
  },
  "timestamp": "2025-09-15T08:35:00Z"
}
```
### **SQLite Payload Example** (terminal offline snapshot + events)
```json
{
  "task": "sync_user_data",
  "user_id": "123",
  "updated_fields": ["current_trip_terminal_id", "balance"],
  "user_snapshot": {
    "user_id": "123",
    "full_name": "Sony Nurdianto",
    "card_id": "XYZ987",
    "balance": 5000,
    "status": "active",
    "current_trip_terminal_id": 2
  },
  "offline_events": [
    {
      "task": "held",
      "terminal_id": "A",
      "hold_amount": 1200,
      "timestamp": "2025-09-15T08:30:00Z"
    },
    {
      "task": "record_checkin",
      "terminal_id": "A",
      "timestamp": "2025-09-15T08:30:00Z",
      "status": "checkin"
    }
  ],
  "timestamp": "2025-09-15T08:35:00Z"
}

```
3️⃣ **Process:**

-   **Sync to Redis:** Update cache supaya future reads cepat
-   **Sync to Terminal SQLite DBs:** Pastikan offline terminals punya snapshot terbaru
-   **Consistency Check:** Bisa optional compare checksums atau versioning supaya semua layer sinkron


4️⃣ **Goal:**

-   Memastikan **data user terbaru** (saldo & current_trip) tersedia di **online** dan **offline mode**
-   Mencegah **inconsistensi saldo atau trip status** saat user check-in / check-out


### 🔎 Observability & Monitoring (Background Tasks)

1.  **East-West Traffic Metrics**
    -   Semua service (Ticketing, Fare, Notification, History, Analytics, Sync) mengirim **metrics** ke **OpenTelemetry Collector**
        -   Metrik bisa berupa request/response latency, error rate, throughput, dsb.
        
2.  **Real-time Monitoring Dashboards**
	-   **Grafana** menampilkan KPI utama:
        
        -   Check-in / Check-out success rate per terminal
            
        -   Response time tiap service / terminal
            
        -   Balance hold operations & failed attempts
            
        -   Queue length & pending tasks di Message Broker (Kafka / RabbitMQ)
            
3.  **Tracing & Logs**
    -   **Jaeger** untuk distributed tracing: pantau flow request antar service
    -   **Loki** untuk centralized log management: debug error, offline events, failed task processing
        
4.  **Goal:**
	 -   Memastikan semua service berjalan normal, performa stabil    
    -   Mendeteksi bottleneck atau masalah di realtime    
    -   Memudahkan troubleshooting jika ada offline/online inconsistency

# 🔄 Fail-Safe Mechanism & Automatic Release System

## ⚠️ Otomatisasi Pelepasan Hold Amount & Pembersihan Data

### 1. **Automatic Hold Amount Release**
#### ⏰ **Scheduled Hold Release Service**
**Background Job**: Berjalan setiap 30 menit untuk melepas hold amount yang tertinggal

**Business Logic**:
```sql
-- Query untuk menemukan hold yang perlu dilepas
SELECT * FROM balance_holds 
WHERE status = 'held' 
AND created_at < NOW() - INTERVAL '24 hours'
AND released_at IS NULL;

-- Process: Release hold amount
UPDATE balance_holds 
SET status = 'released', 
    released_at = NOW(),
    release_reason = 'auto_release_24h_timeout'
WHERE status = 'held' 
AND created_at < NOW() - INTERVAL '24 hours';

-- Kembalikan balance ke user
UPDATE users 
SET balance = balance + hold_amount
FROM balance_holds
WHERE users.id = balance_holds.user_id
AND balance_holds.status = 'released'
AND balance_holds.release_reason = 'auto_release_24h_timeout';
```

**Alasan Release**:
- Hold melebihi 24 jam tanpa checkout
- User tidak menyelesaikan perjalanan
- Prevent balance terkunci indefinitely

### 2. **Orphaned Journey Cleanup**
#### 🧹 **Daily Journey Cleanup Service**
**Background Job**: Berjalan setiap hari jam 02:00 AM

**Cleanup Logic**:
```sql
-- Temukan journey yang tertinggal > 24 jam
UPDATE journey_history 
SET status = 'expired',
    expiry_reason = 'auto_cleanup_24h_no_checkout'
WHERE status = 'active'
AND checkin_time < NOW() - INTERVAL '24 hours';

-- Update user current trip status
UPDATE users 
SET current_trip_terminal_id = NULL
FROM journey_history
WHERE users.id = journey_history.user_id
AND journey_history.status = 'expired'
AND journey_history.expiry_reason = 'auto_cleanup_24h_no_checkout';
```

**Manfaat**:
- Mencegah state inconsistent
- Membersihkan data yang tidak diperlukan
- Memastikan user tidak stuck dalam state "in journey"

### 3. **Notification untuk Automatic Release**
#### 📧 **Auto-Release Notifications**
**User Notification**:
```json
{
  "notification_type": "balance_auto_released",
  "user_id": "123",
  "message": "Dana hold sebesar Rp 1.200 telah dikembalikan ke saldo Anda karena perjalanan tidak diselesaikan dalam 24 jam",
  "amount_released": 1200,
  "release_time": "2024-01-16T08:30:00Z"
}
```

**Admin Alert**:
```json
{
  "alert_type": "auto_release_executed",
  "execution_time": "2024-01-16T02:30:00Z",
  "total_holds_released": 15,
  "total_amount_released": 18000,
  "details": "15 transaksi hold dilepas secara otomatis"
}
```

### 4. **Monitoring & Metrics untuk Auto-Release**
#### 📊 **Auto-Release Monitoring**
**Grafana Dashboard Metrics**:
- Jumlah auto-release per hari
- Total amount yang direlease
- Rata-rata waktu hold sebelum release
- Pattern abandoned journeys

**Alert Rules**:
- High volume auto-release (> 100/day)
- Large amount released (> 1,000,000)
- Seasonal patterns detection

### 5. **Configuration Management**
#### ⚙️ **Configurable Parameters**
**Environment Variables**:
```bash
AUTO_RELEASE_TIMEOUT_HOURS=24
CLEANUP_SCHEDULE="0 2 * * *"
MAX_HOLD_AMOUNT=50000
MIN_BALANCE_THRESHOLD=5000
```

**Runtime Configuration**:
- Timeout duration dapat diadjust
- Threshold values dapat dimodifikasi
- Enable/disable feature per terminal

## 🎯 **Manfaat Fail-Safe Mechanism:**

### 1. **User Experience**
- ✅ Tidak ada dana terkunci selamanya
- ✅ Transparent communication ke user
- ✅ Automatic problem resolution

### 2. **Business Protection**  
- ✅ Mencegah financial discrepancies
- ✅ Maintain accurate accounting
- ✅ Reduce customer service complaints

### 3. **System Health**
- ✅ Automated cleanup
- ✅ Preventive maintenance
- ✅ Data consistency guarantees

### 4. **Operational Efficiency**
- ✅ Reduced manual intervention
- ✅ Proactive issue resolution
- ✅ Comprehensive monitoring

## 🔄 **Integration dengan Existing Flow:**

### **Ditambahkan ke Flow Online:**
#### Step 8: Background Processes (Additional)
1. **Scheduled Hold Release** - Setiap 30 menit
2. **Journey Cleanup** - Setiap hari jam 02:00 AM  
3. **Notification Service** - Kirim auto-release alerts
4. **Monitoring** - Track metrics dan generate reports

### **Ditambahkan ke Database Schema:**
```sql
-- Tambah column untuk tracking release reason
ALTER TABLE balance_holds 
ADD COLUMN release_reason VARCHAR(255) NULL;

-- Tambah column untuk expiry reason
ALTER TABLE journey_history
ADD COLUMN expiry_reason VARCHAR(255) NULL;
```

## 📋 **Complete Fail-Safe Coverage:**

| **Scenario** | **Detection** | **Action** | **Recovery** |
|-------------|---------------|------------|--------------|
| Hold > 24h | Time-based check | Release hold | Balance returned |
| No checkout | Journey timeout | Mark expired | State cleaned up |
| System crash | Monitoring alerts | Auto-restart | Service recovered |
| Data corrupt | Consistency check | Repair data | Integrity restored |


---

## 🏁 Flow 2: Check-Out Process (Online)

### Step 1: User Interaction  
1. **User** tap kartu NFC di **Gate Device Exit Terminal ** (tujuan akhir)
2. Gate Device capture card data dan buat HTTP request ke **CheckOut endpoint**

### Step 2: Request Routing (Same as Check-in)
1. **Terminal Akhir** → **Istio Ingress** → **Service Mesh** → **Ticketing Service**
2. **North-South observability** ke OpenTelemetry

### Step 3: User Data Retrieval & Validation
**Ticketing Service** menggunakan same priority strategy:
1. **Redis** → **PostgreSQL** → **SQLite** (fallback)

**Enhanced Validation for Checkout**:
1. **User Existence & Status**: Same as check-in
2. **Active Journey Validation**:
   - Check `current_trip_terminal_id` IS NOT NULL
   - Jika NULL → User belum check-in → Response "No Active Journey" → Gate **CLOSED**
   - Display: "Anda belum check-in, silakan check-in terlebih dahulu"

### Step 4: Fare Calculation Request
**Ticketing Service** kirim request ke **Fare Service** dengan endpoint khusus checkout:
```json
{
  "action": "checkout",
  "user_id": "123",
  "checkin_terminal": "A", // from current_trip_terminal_id
  "checkout_terminal": "E",
  "timestamp": "2024-01-15T09:15:00Z"
}
```

### Step 5: Fare Service Processing
**Fare Service** process checkout:

1. **Fare Calculation**:
   - Lookup fare matrix: A → E = 1200
   - Current held amount: 1200
   - Actual fare needed: 1200
   - Balance adjustment: 0 (sudah pas)

2. **Different Scenarios**:
   ```
   Scenario 1: A → B (fare 500)
   - Held: 1200, Actual: 500
   - Refund: 700 back to balance
   
   Scenario 2: A → E (fare 1200)  
   - Held: 1200, Actual: 1200
   - No refund needed
   
   Scenario 3: Edge case - insufficient (rare)
   - Held: 1200, but rule changed to 1500
   - Additional deduction from balance
   ```

3. **Balance Update**:
   - Release hold amount (1200)
   - Charge actual fare (1200)  
   - Final balance = original_balance - actual_fare
   - Update PostgreSQL with final transaction

### Step 6: Gate Control & Response
1. **Success Response**:
   - Ticketing Service receive confirmation dari Fare Service
   - Send **"OPEN GATE"** command ke Terminal Akhir
   - Gate terbuka, user bisa exit

2. **Update User State**:
   - Set `current_trip_terminal_id = NULL` (journey completed)

### Step 7: Async Task Distribution

**Ticketing Service** publish event ke **Message Broker** untuk 5 async tasks:

#### Task 1: **Fare Service** - Balance Release & Final Charge

1️⃣ Data Structure: Update `balance_holds` table
```sql
UPDATE balance_holds 
SET status = 'charged', 
    released_at = NOW(),
    final_fare = 1200
WHERE user_id = 123 AND status = 'held';
```
2️⃣ Event Payload (Message Broker)
```json
{
  "task": "release_and_charge",
  "user_id": "123",
  "checkin_terminal_id": "A",
  "checkout_terminal_id": "E", 
  "held_amount": 1200,
  "actual_fare": 1200,
  "refund_amount": 0,
  "timestamp": "2024-01-15T09:15:00Z"
}
```
**Fare Service** update PostgreSQL:

-   Release held amount (1200)
-   Charge actual fare (1200)
-   Final balance = original_balance - actual_fare + refund_amount
-   Mark transaction sebagai "charged"


#### Task 2: **Notification Service** - CheckOut Alert
1️⃣ Data Structure: Insert ke `notifications` table (same structure as CheckIn)
```sql
notifications {
    id: PRIMARY KEY
    user_id: INT REFERENCES users(id)
    type: ENUM('checkin','checkout','fare_update')
    status ENUM('pending','sent','failed')
    created_at: TIMESTAMP
    sent_at: TIMESTAMP NULL
    message TEXT
}
```
2️⃣ Event Payload (Message Broker)
```json
{
  "task": "send_checkout_notification",
  "user_id": "123",
  "checkin_terminal_id": "A",
  "checkout_terminal_id": "E",
  "fare_charged": 1200,
  "journey_duration_minutes": 45,
  "final_balance": 3800
}
```
Send email: "Perjalanan selesai A → E pada 09:15. Tarif: Rp 1.200. Sisa saldo: Rp 3.800. Durasi: 45 menit."

#### Task 3: **History Service** - Journey Record Completiond
1️⃣ Data Structure: Update existing `journey_history` record
```sql
UPDATE journey_history 
SET checkout_terminal_id = 'E',
    checkout_time = '2024-01-15T09:15:00Z',
    fare_charged = 1200,
    status = 'completed'
WHERE user_id = 123 AND status = 'active';
```
2️⃣ Event Payload (Message Broker)
```json
{
  "task": "complete_journey_record",
  "user_id": "123",
  "checkout_terminal_id": "E",
  "checkout_time": "2024-01-15T09:15:00Z",
  "fare_charged": 1200,
  "status": "checkout"
}
```
**History Service** update existing journey record:
1. Find active journey: `WHERE user_id = 123 AND status = 'active'`  
-   `checkout_terminal_id`: E
-   `checkout_time`: 2024-01-15T09:15:00Z
-   `fare_charged`: 1200
-   `status`: 'completed'
-   Calculate `journey_duration`: checkout_time - checkin_time


#### Task 4: **Analytics Service** - CheckOut Usage Data
1️⃣ **Data Structure: Insert ke `analytics_events` table** (same structure as CheckIn)
```sql
analytics_events {
    id: PRIMARY KEY
    event_type: ENUM('checkin','checkout','fare_update')
    terminal_id: INT REFERENCES terminals(id)
    user_id: INT NULL REFERENCES users(id)
    user_demographics JSONB NULL
    event_timestamp: TIMESTAMP
    created_at: TIMESTAMP DEFAULT now()
}
```
2️⃣ Event Payload (Message Broker)
```json
{
  "task": "record_checkout_analytics",
  "event_type": "checkout",
  "terminal_id": "E",
  "user_id": "123",
  "timestamp": "2024-01-15T09:15:00Z",
  "journey_data": {
    "origin_terminal": "A",
    "destination_terminal": "E",
    "fare_amount": 1200,
    "journey_duration_minutes": 45
  },
  "user_demographics": {
      "age_group": "25-34",
      "gender": "male", 
      "membership": "gold"
  }
}
```
**Analytics Service** insert ke `analytics_events` untuk:

-   CheckOut event tracking
-   Route popularity analysis (A → E)
-   Revenue calculation
-   Journey duration patterns
-   Peak hour analysis


#### Task 5: **Sync Data Service** - Data Consistency (CheckOut)

1️⃣ **Data Structure: `offline_tasks` table untuk SQLite Offline mode**
```sql
offline_tasks {
    id: PRIMARY KEY,
    service_name: VARCHAR,
    action: VARCHAR, 
    payload: JSONB,
    status: ENUM('pending','processed','failed'),
    created_at: TIMESTAMP,
    processed_at: TIMESTAMP NULL
}
```
-   **users** table (PostgreSQL)
-   **Redis Cache** (fast lookup)
-   **SQLite snapshots** di tiap terminal (offline mode)

**Redis Payload Example** (fast lookup cache)
```json
{
  "task": "sync_checkout_data",
  "user_id": "123",
  "updated_fields": ["current_trip_terminal_id", "balance"],
  "user_snapshot": {
    "user_id": "123",
    "full_name": "Sony Nurdianto",
    "card_id": "XYZ987",
    "balance": 3800,
    "status": "active",
    "current_trip_terminal_id": null
  },
  "timestamp": "2024-01-15T09:20:00Z"
}
```
**SQLite Payload Example** (terminal offline snapshot + events)
```json
{
  "task": "sync_checkout_data",
  "user_id": "123", 
  "updated_fields": ["current_trip_terminal_id", "balance"],
  "user_snapshot": {
    "user_id": "123",
    "full_name": "Sony Nurdianto", 
    "card_id": "XYZ987",
    "balance": 3800,
    "status": "active",
    "current_trip_terminal_id": null
  },
  "offline_events": [
    {
      "task": "release_and_charge",
      "checkin_terminal_id": "A",
      "checkout_terminal_id": "E",
      "actual_fare": 1200,
      "timestamp": "2024-01-15T09:15:00Z"
    },
    {
      "task": "complete_journey_record", 
      "checkout_terminal_id": "E",
      "timestamp": "2024-01-15T09:15:00Z",
      "status": "checkout"
    }
  ],
  "timestamp": "2024-01-15T09:20:00Z"
}
```
3️⃣ **Process:**

-   **Sync to Redis:** Update cache supaya future reads cepat dengan final balance dan clear current_trip
-   **Sync to Terminal SQLite DBs:** Pastikan offline terminals punya snapshot terbaru dengan journey completed
-   **Consistency Check:** Bisa optional compare checksums atau versioning supaya semua layer sinkron

4️⃣ **Goal:**

-   Memastikan **final balance dan completed journey status** tersedia di **online** dan **offline mode**
-   Mencegah **double checkout** atau **orphaned trip status** saat user check-out ulang

---


# E-Ticketing System - Offline CheckIn & CheckOut Flow

## 📋 Offline System Overview

Ketika terminal kehilangan koneksi internet, sistem beralih ke **Offline Mode** menggunakan local SQLite database. Semua transaksi disimpan sebagai **pending tasks** dan akan diproses saat connectivity restored.

### Offline Data Structures

#### Local SQLite Database (per terminal)

```sql
-- User snapshots for offline validation
users_snapshot {
    id: PRIMARY KEY
    full_name: VARCHAR
    card_id: VARCHAR UNIQUE
    balance: DECIMAL
    status: ENUM('active', 'blocked')
    current_trip_terminal_id: INT NULL
    last_sync_timestamp: TIMESTAMP
}

-- Local fare matrix
fare_matrix {
    origin_terminal_id: INT
    destination_terminal_id: INT 
    price: DECIMAL
    last_sync_timestamp: TIMESTAMP
}

-- Pending tasks queue
pending_tasks {
    task_id: VARCHAR PRIMARY KEY
    task_type: ENUM('checkin','checkout')
    user_id: INT
    terminal_id: INT
    checkin_terminal_id: INT NULL  -- for checkout tasks
    checkout_terminal_id: INT NULL -- for checkout tasks
    estimated_fare: DECIMAL NULL
    payload: JSONB
    timestamp: TIMESTAMP
    status: ENUM('pending','processing','completed','failed')
}

```

----------

## 🔄 Flow 3: CheckIn Process (Offline)

### Step 1: User Interaction

1.  **User** tap kartu NFC atau mobile dengan NFC feature ke **Gate Device Terminal A**
2.  Gate Device detect **network connectivity lost**
3.  Gate Device capture card data dan route ke **Local SQLite Processing**

### Step 2: Offline Mode Detection & Local Routing

1.  Gate Device detect **network timeout** dari Istio Ingress
2.  **Automatic Fallback**: Switch ke **Local Processing Mode**
3.  **Local Observability**: Store metrics di local buffer untuk sync nanti
4.  Display indicator: "MODE OFFLINE - Transaksi akan diproses saat sistem normal"

### Step 3: User Data Retrieval (Offline Priority Strategy)

**Local Terminal Processing** menggunakan offline priority:

1.  **Only Priority**: Query **Local SQLite users_snapshot**
    -   Jika user data ditemukan → proceed to validation
    -   Data terakhir sync dari Redis/PostgreSQL sebelum offline

### Step 4: User Validation (Offline CheckIn)

**Local Processing** validate user data:

1.  **User Existence Check**:
    
    -   Jika user tidak ditemukan di local snapshot → Response "Not Registered" → Gate tetap **CLOSED**
    -   Display message: "Kartu belum terdaftar atau data belum sync, coba lagi nanti"
2.  **User Status Check**:
    
    -   Jika status = 'blocked' → Response "Unauthorized" → Gate tetap **CLOSED**
    -   Display message: "Kartu diblokir, hubungi customer service"
3.  **Active Journey Check**:
    
    -   Jika `current_trip_terminal_id` IS NOT NULL → User masih dalam perjalanan
    -   Response "Already Checked In" → Gate tetap **CLOSED**
    -   Display message: "Anda sudah check-in, lakukan check-out terlebih dahulu"

### Step 5: Balance Validation (Offline)

1.  **Local Balance Check**:
    
    -   Check `user_snapshot.balance >= minimum_fare` (500)
    -   Jika insufficient → Response "Insufficient Balance" → Gate tetap **CLOSED**
    -   Display balance terakhir yang di-sync
2.  **Conservative Hold Strategy**:
    
    -   Hold maximum fare (1200) dari local balance snapshot
    -   **Risk Management**: Optimistic deduction dengan reconciliation nanti

### Step 6: Gate Control & Response (Offline)

1.  **Success Response**:
    
    -   Local Processing kirim command **"OPEN GATE"** ke Terminal CheckIn
    -   Gate terbuka selama 5 detik
    -   Display: "Check-in berhasil (offline). Saldo akan diverifikasi saat sistem online"
2.  **Update Local State**:
    
    -   Update `users_snapshot.current_trip_terminal_id = A`
    -   Update `users_snapshot.balance = balance - 1200` (optimistic hold)

### Step 7: Local Task Queue Creation

**Local Processing** create 5 pending tasks di SQLite queue:

#### Task 1: **Fare Service** - Offline Balance Hold

1️⃣ Pending Task Creation

```sql
INSERT INTO pending_tasks (
    task_id, task_type, user_id, terminal_id,
    estimated_fare, payload, timestamp, status
) VALUES (
    'offline-hold-uuid-123', 'checkin', 123, 'A',
    1200, 
    '{"task": "held", "hold_amount": 1200, "terminal_id": "A"}',
    '2024-01-15T08:30:00Z', 'pending'
);

```

2️⃣ Local Optimistic Update

-   Deduct 1200 dari local `users_snapshot.balance`
-   Mark sebagai 'pending_hold' untuk reconciliation

#### Task 2: **Notification Service** - Offline CheckIn Alert

1️⃣ Pending Task Creation

```sql
INSERT INTO pending_tasks (
    task_id, task_type, user_id, terminal_id,
    payload, timestamp, status
) VALUES (
    'offline-notif-uuid-124', 'checkin', 123, 'A',
    '{"task": "send_checkin_notification", "hold_amount": 1200, "mode": "offline"}',
    '2024-01-15T08:30:00Z', 'pending' 
);

```

2️⃣ Local Display Message

-   Show success dengan disclaimer offline mode
-   "Notifikasi email akan dikirim saat sistem online"

#### Task 3: **History Service** - Offline Journey Record

1️⃣ Pending Task Creation

```sql
INSERT INTO pending_tasks (
    task_id, task_type, user_id, terminal_id,
    payload, timestamp, status
) VALUES (
    'offline-history-uuid-125', 'checkin', 123, 'A',
    '{"task": "record_checkin", "status": "checkin", "mode": "offline"}',
    '2024-01-15T08:30:00Z', 'pending'
);

```

2️⃣ Local Temporary Record

-   Store basic journey info di local untuk reference

#### Task 4: **Analytics Service** - Offline Usage Data

1️⃣ Pending Task Creation

```sql
INSERT INTO pending_tasks (
    task_id, task_type, user_id, terminal_id,
    payload, timestamp, status
) VALUES (
    'offline-analytics-uuid-126', 'checkin', 123, 'A',
    '{"task": "record_analytics", "event_type": "checkin", "mode": "offline"}',
    '2024-01-15T08:30:00Z', 'pending'
);

```

2️⃣ Local Metrics Buffer

-   Store analytics data di local buffer untuk batch sync

#### Task 5: **Sync Data Service** - Offline Data Reconciliation

1️⃣ Pending Task Creation

```sql
INSERT INTO pending_tasks (
    task_id, task_type, user_id, terminal_id,
    payload, timestamp, status
) VALUES (
    'offline-sync-uuid-127', 'checkin', 123, 'A', 
    '{"task": "sync_user_data", "updated_fields": ["current_trip_terminal_id", "balance"], "mode": "offline"}',
    '2024-01-15T08:30:00Z', 'pending'
);

```

2️⃣ Local State Tracking

-   Mark user data sebagai 'needs_sync'
-   Track offline changes untuk reconciliation

----------

## 🏁 Flow 4: CheckOut Process (Offline)

### Step 1: User Interaction

1.  **User** tap kartu NFC atau mobile dengan NFC feature ke **Gate Device Terminal E**
2.  Gate Device detect **network connectivity lost**
3.  Gate Device capture card data dan route ke **Local SQLite Processing**

### Step 2: Offline Mode Detection & Local Routing

1.  Gate Device detect **network timeout** dari Istio Ingress
2.  **Automatic Fallback**: Switch ke **Local Processing Mode**
3.  **Local Observability**: Store metrics di local buffer untuk sync nanti
4.  Display indicator: "MODE OFFLINE - Transaksi akan diproses saat sistem normal"

### Step 3: User Data Retrieval (Offline Priority Strategy)

**Local Terminal Processing** menggunakan offline priority:

1.  **Only Priority**: Query **Local SQLite users_snapshot**
    -   Jika user data ditemukan → proceed to validation
    -   Data terakhir sync dari Redis/PostgreSQL sebelum offline

### Step 4: User Validation (Offline CheckOut)

**Local Processing** validate user data:

1.  **User Existence Check**:
    
    -   Jika user tidak ditemukan di local snapshot → Response "Not Registered" → Gate tetap **CLOSED**
    -   Display message: "Kartu belum terdaftar atau data belum sync, coba lagi nanti"
2.  **User Status Check**:
    
    -   Jika status = 'blocked' → Response "Unauthorized" → Gate tetap **CLOSED**
    -   Display message: "Kartu diblokir, hubungi customer service"
3.  **Active Journey Check** (CheckOut Specific):
    
    -   Jika `current_trip_terminal_id` IS NULL → User belum check-in
    -   Response "No Active Journey" → Gate tetap **CLOSED**
    -   Display message: "Anda belum check-in, silakan check-in terlebih dahulu"

### Step 5: CheckOut Journey Validation & Fare Calculation (Offline)

1.  **Local Active Journey Check**:
    
    -   Confirm `current_trip_terminal_id` IS NOT NULL di local snapshot
    -   Get checkin terminal dari `users_snapshot.current_trip_terminal_id`
2.  **Local Fare Calculation**:
    
    -   Query local `fare_matrix`: checkin_terminal (A) → checkout_terminal (E) = 1200
    -   **Business Decision**: Use local fare matrix (might be stale)
    -   **Risk Management**: Allow checkout dengan fare reconciliation nanti

### Step 6: Gate Control & Response (Offline)

1.  **Success Response**:
    
    -   Local Processing kirim command **"OPEN GATE"** ke Terminal CheckOut
    -   Gate terbuka selama 5 detik
    -   Display: "Check-out berhasil (offline). Tarif Rp 1.200 akan diverifikasi saat sistem online"
2.  **Update Local State**:
    
    -   Set `users_snapshot.current_trip_terminal_id = NULL` (journey completed)
    -   Optimistic balance calculation (bisa salah, akan direconcile)

### Step 7: Local Task Queue Creation (CheckOut)

**Local Processing** create 5 pending tasks di SQLite queue:

#### Task 1: **Fare Service** - Offline Balance Release & Charge

1️⃣ Pending Task Creation

```sql
INSERT INTO pending_tasks (
    task_id, task_type, user_id, checkin_terminal_id, checkout_terminal_id,
    estimated_fare, payload, timestamp, status
) VALUES (
    'offline-fare-uuid-201', 'checkout', 123, 'A', 'E',
    1200,
    '{"task": "release_and_charge", "actual_fare": 1200, "mode": "offline"}', 
    '2024-01-15T09:15:00Z', 'pending'
);

```

2️⃣ Local Optimistic Processing

-   Release held amount (1200) di local snapshot
-   Charge actual fare (1200) - zero adjustment in this case

#### Task 2: **Notification Service** - Offline CheckOut Alert

1️⃣ Pending Task Creation

```sql
INSERT INTO pending_tasks (
    task_id, task_type, user_id, checkin_terminal_id, checkout_terminal_id,
    estimated_fare, payload, timestamp, status  
) VALUES (
    'offline-notif-uuid-202', 'checkout', 123, 'A', 'E',
    1200,
    '{"task": "send_checkout_notification", "fare_charged": 1200, "mode": "offline"}',
    '2024-01-15T09:15:00Z', 'pending'
);

```

2️⃣ Local Display Message

-   "Check-out berhasil. Email konfirmasi akan dikirim saat sistem online"

#### Task 3: **History Service** - Offline Journey Record Completion

1️⃣ Pending Task Creation

```sql
INSERT INTO pending_tasks (
    task_id, task_type, user_id, checkin_terminal_id, checkout_terminal_id,
    estimated_fare, payload, timestamp, status
) VALUES (
    'offline-history-uuid-203', 'checkout', 123, 'A', 'E', 
    1200,
    '{"task": "complete_journey_record", "status": "checkout", "mode": "offline"}',
    '2024-01-15T09:15:00Z', 'pending'
);

```

2️⃣ Local Journey Completion

-   Mark local journey sebagai 'completed_offline' untuk tracking

#### Task 4: **Analytics Service** - Offline CheckOut Usage Data

1️⃣ Pending Task Creation

```sql
INSERT INTO pending_tasks (
    task_id, task_type, user_id, checkin_terminal_id, checkout_terminal_id,
    estimated_fare, payload, timestamp, status
) VALUES (
    'offline-analytics-uuid-204', 'checkout', 123, 'A', 'E',
    1200, 
    '{"task": "record_checkout_analytics", "event_type": "checkout", "mode": "offline"}',
    '2024-01-15T09:15:00Z', 'pending'
);

```

2️⃣ Local Analytics Buffer

-   Store route analytics (A → E) di local untuk batch processing

#### Task 5: **Sync Data Service** - Offline Data Reconciliation (CheckOut)

1️⃣ Pending Task Creation

```sql
INSERT INTO pending_tasks (
    task_id, task_type, user_id, checkin_terminal_id, checkout_terminal_id,
    estimated_fare, payload, timestamp, status
) VALUES (
    'offline-sync-uuid-205', 'checkout', 123, 'A', 'E',
    1200,
    '{"task": "sync_checkout_data", "updated_fields": ["current_trip_terminal_id", "balance"], "mode": "offline"}',
    '2024-01-15T09:15:00Z', 'pending'
);

```

2️⃣ Local State Final Update

-   Clear `current_trip_terminal_id` dari local snapshot
-   Mark sebagai 'needs_reconciliation' untuk online sync

----------

## 🔄 Flow 5: Online Recovery & Reconciliation

### Step 1: Connectivity Restored Detection

1.  **Network Monitor** detect koneksi kembali normal
2.  **Local Processing** switch back ke **Online Mode**
3.  **Sync Data Service** start reconciliation process

### Step 2: Pending Tasks Retrieval

**Sync Data Service** query semua terminals untuk pending tasks:

```sql
SELECT * FROM pending_tasks 
WHERE status = 'pending'
ORDER BY timestamp ASC; -- Process chronologically  

```

### Step 3: Task Processing & Distribution

**Sync Data Service** publish pending tasks ke **Message Broker** sesuai priority:

#### Priority 1: CheckOut Tasks (Free stuck users)

```json
{
  "reconciliation_batch": "checkout_priority",
  "tasks": [
    {
      "original_task_id": "offline-fare-uuid-201",
      "task_type": "checkout",
      "user_id": "123",
      "checkin_terminal_id": "A", 
      "checkout_terminal_id": "E",
      "offline_timestamp": "2024-01-15T09:15:00Z",
      "estimated_fare": 1200
    }
  ]
}

```

#### Priority 2: CheckIn Tasks

```json
{
  "reconciliation_batch": "checkin_priority",
  "tasks": [
    {
      "original_task_id": "offline-hold-uuid-123", 
      "task_type": "checkin",
      "user_id": "123",
      "terminal_id": "A",
      "offline_timestamp": "2024-01-15T08:30:00Z",
      "estimated_hold": 1200
    }
  ]
}

```

### Step 4: Service-Specific Reconciliation

#### **Fare Service** Reconciliation

1.  **CheckOut Tasks**:
    
    -   Validate actual fare vs estimated fare
    -   Process balance adjustment if needed
    -   Handle fare matrix changes during offline period
2.  **CheckIn Tasks**:
    
    -   Validate current balance vs offline hold
    -   Process actual balance hold
    -   Handle insufficient balance scenarios

#### **History Service** Reconciliation

1.  Insert/Update journey records dengan flag 'offline_recovered'
2.  Calculate accurate journey duration
3.  Mark offline periods in journey history

#### **Notification Service** Reconciliation

1.  Send delayed notifications dengan offline context
2.  Include reconciliation results if applicable
3.  Send bulk summary untuk multiple offline transactions

### Step 5: Conflict Resolution & Data Consistency

#### **Conflict Scenarios**:

1.  **Double CheckIn**: User check-in offline di A, then online di B
    
    -   **Resolution**: Honor first chronologically, reject second
2.  **Insufficient Balance**: Offline deduction exceeded actual balance
    
    -   **Resolution**: Allow temporary negative, flag for customer service
3.  **Fare Changes**: Offline used stale fare matrix
    
    -   **Resolution**: Honor offline fare for fairness

#### **Data Sync Priority**:

1.  **Central Truth**: PostgreSQL reconciled data
2.  **Cache Update**: Refresh Redis dengan latest state
3.  **Local Sync**: Update all SQLite snapshots
4.  **Audit Trail**: Log all reconciliation activities

### Step 6: Monitoring & Alerting (Reconciliation)

#### **Real-time Dashboards** (Grafana)

1.  **Reconciliation Metrics**:
    
    -   Pending tasks count per terminal
    -   Processing success rate
    -   Reconciliation duration
    -   Conflict resolution count
2.  **Business Impact**:
    
    -   Revenue impact dari fare changes
    -   User experience metrics
    -   Offline duration analysis

#### **Critical Alerts**

1.  **High pending count** (> 100 tasks)
2.  **Reconciliation failures** (> 5%)
3.  **Extended offline periods** (> 1 hour)
4.  **Balance discrepancies** > threshold (1000)

----------

## 📊 Offline vs Online Flow Comparison
| Aspect             | Online Flow                           | Offline Flow                          |
|-------------------|--------------------------------------|--------------------------------------|
| Data Source        | Redis → PostgreSQL → SQLite           | SQLite Only                           |
| Processing         | Real-time via Service Mesh            | Local optimistic + queue              |
| Gate Response      | Immediate with live data              | Immediate with stale data             |
| Task Distribution  | Message Broker (async)                | Local queue (pending)                 |
| Data Consistency   | Immediate sync                        | Eventual consistency                  |
| Error Handling     | Real-time validation                  | Reconciliation-based                  |
| User Experience    | Full feature set                      | Limited with disclaimers              |
| Business Risk      | Low (real-time validation)            | Medium (reconciliation needed)       |

---

## 🛠 Technology Stack Proposal

1.  ### **Communication Layer :**
    
    -   ###**gRPC Fast , Eficient and Light Communication** 
	    -  Menggunakan Protokol Biner (Protobuf) untuk serialisasi data, yang menghasilkan payload yang lebih kecil dan parsing yang lebih cepat dibandingkan REST/JSON. **Efisiensi ini sangat menghemat bandwidth dan mengurangi latency.**
	    - Mendukung pola komunikasi Streaming (Server Streaming, Client Streaming, dan **Bidirectional Streaming**). **Kemampuan Bidirectional Streaming sangat cocok untuk fitur real-time seperti notifikasi live order atau tracking lokasi ** yang cocok untuk High Frequency Transaction seperti (checkin/checkout).
	    -  Pendekatan **Contract-First** dengan Protobuf (`*.proto` files) **menjadi sumber kebenaran tunggal** yang memudahkan kolaborasi, menjaga konsistensi, dan meningkatkan maintainability antar tim/service yang berbeda.
	    -  **Memiliki fitur keamanan (Security) built-in yang kuat dengan menggunakan HTTP/2 sebagai transport dan mendukung TLS/SSL encryption secara native,** sehingga komunikasi antar service terjamin kerahasiaannya.
	    -  **Memungkinkan generate code secara otomatis** (boilerplate code) untuk berbagai bahasa pemrograman (Go, Java, Python, dll) dari file `.proto`, **yang mempercepat development dan mengurangi human error.**

        
    -   ### **REST API (Representational State Transfer API)**
	    - **Compatibility & Wider Accessibility**: REST menggunakan standar HTTP/1.1 dan JSON yang universal, sehingga **sangat ideal untuk integrasi dengan pihak eksternal** seperti fintech, payment gateway, regulator, atau client aplikasi mobile/web front-end. **Pihak ketiga tidak memerlukan library atau tool khusus untuk berkomunikasi dengan sistem kita.**
    
		-   **Broader Ecosystem & Simplicity**: Didukung oleh hampir semua bahasa pemrograman dan framework tanpa dependensi tambahan yang kompleks. **Ini mengurangi barrier untuk adoption bagi partner yang mungkin tidak familiar dengan RPC/gRPC.**
    
		-   **Simpler Debugging & Operational Visibility**: **Operasi CRUD (GET, POST, PUT, DELETE) yang sederhana dan align dengan semantic HTTP.** Setiap request dapat dengan mudah diinspeksi, di-test, dan di-debug menggunakan tool sederhana seperti **browser developer tools, Postman, atau cURL.** Hal ini sangat disukai oleh ops team dan developer untuk troubleshooting.
    
		-   **Suitable for Low-Frequency Use Cases**: **Overhead HTTP/1.1 dan JSON yang lebih besar membuatnya kurang efisien untuk komunikasi internal high-frequency.** Namun, justru karena itu REST sangat cocok untuk **API yang tidak memerlukan throughput sangat tinggi atau real-time streaming**, seperti API untuk dashboard admin, laporan, atau proses yang berjalan cron-job.
	
	- ### **Message Broker (Kafka Or RabbitMQ)**
		- **RabbitMQ: Reliable, Flexible, and Easy-to-Integrate Messaging**
			-  **Queue-based Messaging** , RabbitMQ menggunakan model antrian tradisional, sehingga cocok untuk **transaksi checkin/checkout** yang harus diproses **satu per satu dengan urutan terjamin**.
    
			-   **Flexible Routing (Exchange Types)** , mendukung **direct, topic, fanout, dan headers exchange**, sehingga message bisa diarahkan sesuai kebutuhan (misal hanya ke payment service atau broadcast ke monitoring + analytics).
    
			-   **Acknowledgement & Retry** , mendukung **manual ACK/NACK** dan **dead-letter queues (DLQ)**, memastikan tidak ada transaksi yang hilang atau terabaikan meski ada error.
    
			-   **Wide Protocol Support** , tidak hanya AMQP, RabbitMQ juga mendukung **STOMP, MQTT, WebSocket**, sehingga mudah diintegrasikan dengan berbagai perangkat keras atau sistem lama.
    
			-   **Simplicity & Ecosystem** , lebih ringan dipelajari, dokumentasi lengkap, dan banyak library siap pakai di berbagai bahasa. Cocok kalau perusahaan sudah terbiasa dengan pola **traditional message queue**. 
			
		- **Apache Kafka High-Throughput, Scalable, and Real-Time Event Streaming**
			- **Log-based Architecture**: **Menyimpan data sebagai log yang append-only dan terurut (ordered sequence of events).** Setiap event (e.g., "CheckinAttempted", "CheckoutCompleted") **dapat dikonsumsi berkali-kali oleh banyak consumer group yang independen.** Ini membangun **single source of truth** yang ideal untuk audit trail, real-time analytics, dan stream processing.
    
			-   **High Throughput & Low Latency**: **Dirancang untuk menangani volume data yang sangat masif** (jutaan event per detik) dengan latensi minimal. **Kinerja ini sangat critical untuk menghandle puluhan ribu gate yang melakukan checkin/checkout secara bersamaan pada jam sibuk.**
    
			-   **Horizontal Scalability & Fault Tolerance**: **Mudah discale secara horizontal** dengan menambahkan broker baru. Data yang dipartisi (**partitioned**) dan direplikasi (**replicated**) di across broker menjamin **tidak ada single point of failure** dan ketersediaan sistem tetap tinggi.
    
			-   **Event Sourcing & Streaming Patterns**: **Lebih dari sekadar queue, Kafka adalah platform event streaming.** Setiap perubahan status (state) dalam sistem direpresentasikan sebagai event yang immutable. Pola ini **sangat powerful untuk membangun sistem event-driven** di mana services seperti Payment, Notification, dan Analytics dapat bereaksi terhadap event yang sama secara real-time dan independen.
    
			-   **Durability & Configurable Retention**: **Event tidak dihapus setelah dikonsumsi, tetapi disimpan secara persistent di disk berdasarkan durasi atau ukuran yang ditentukan.** Hal ini memungkinkan **reprocessing data dari masa lalu** untuk keperluan debugging, membangun ulang dataset, atau compliance.
    
			-   **Rich Ecosystem**: Didukung oleh tools seperti **Kafka Connect** untuk integrasi data dengan database (source/sink) dan **Kafka Streams** (atau **ksqlDB**) untuk memproses aliran data secara real-time. **Ini membuat Kafka menjadi tulang punggung (backbone) data untuk seluruh organisasi.**

		### **Kafka vs. RabbitMQ: Ringkasan Pros and Cons**

		#### **Apache Kafka: The Event Streaming Powerhouse**

		-   **👍 Pros (Kekuatan):**
		    
		    -   **Throughput Luar Biasa:** Juara dalam menangani **volume data yang sangat masif** (jutaan event/detik). Ideal jika jumlah penumpang sangat besar.
		        
		    -   **Durable Storage:** Event disimpan lama (bahkan tahunan), jadi bisa diproses ulang kapan saja. Sempurna untuk **audit trail** dan **data analytics**.
		        
		    -   **Stream Processing:** Bisa memproses data secara real-time langsung di stream-nya (pakai Kafka Streams/ksqlDB). Misalnya, menghitung jumlah penumpang per rute secara live.
		        
		    -   **Pola Pub/Sub Skala Besar:** Satu event bisa dikirim ke _consumer_ yang jumlahnya sangat banyak tanpa mempengaruhi performance.
		        
		-   **👎 Cons (Kelemahan untuk Konteks Ini):**
		    
		    -   **Kompleksitas Operasional:** Setup, konfigurasi, dan maintenance cluster Kafka itu **rumit**. Butuh tim DevOps yang dedicated.
		        
		    -   **"Overkill" untuk Task-Based Messaging:** Jika kebutuhan utamanya adalah mengantrikan dan mendistribusikan _task_ (e.g., "proses pembayaran ini", "buka gate itu"), Kafka bisa terasa berat dan kurang fleksibel.
		        
		    -   **Kurang Cocok untuk Request/Reply:** Pola komunikasi yang dibutuhkan untuk meminta dan mendapatkan jawaban segera (seperti request buka gate) bukanlah kekuatan Kafka.
		        
		    -   **Learning Curve:** Konsep seperti partition, offset, consumer group, dan replication bisa jadi tantangan bagi developer yang belum berpengalaman.
		        

		#### **RabbitMQ: The Reliable Message Broker**

		-   **👍 Pros (Kekuatan untuk Use Case Ini):**
		    
		    -   **Sangat Cocok untuk Task Distribution:** Filosofi antrian (queue)-nya **sangat align** dengan kebutuhan memproses perintah "checkin" dan "checkout" sebagai _task_ yang harus dikerjakan sekali dan terjamin.
		        
		    -   **Flexible Routing:**  **Exchange types** (direct, topic, fanout) sangat powerful untuk mengarahkan pesan ke service yang tepat. Misal: pesan "checkin.error" bisa di-route ke queue untuk monitoring dan tim ops.
		        
		    -   **Manajemen Konsumen yang Mudah:** Fitur **ACK/NACK** dan **Dead Letter Queue** sangat intuitif untuk mengelola pesan yang gagal diproses.
		        
		    -   **Operasional Lebih Sederhana:** Lebih mudah di-setup, dikonfigurasi, dan dimonitor. Cocok untuk tim yang ingin fokus pada _business logic_ tanpa pusing mengelola infrastruktur messaging yang terlalu kompleks.
		        
		    -   **Protocol Support:** Dukung AMQP, MQTT, dll. Ini bisa berguna jika di masa depan perlu integrasi dengan perangkat IoT atau sistem lain.
		        
		-   **👎 Cons (Kelemahan):**
		    
		    -   **Performance:** Meski cepat, **throughput maksimalnya tidak se-Kafka**. Bisa jadi bottleneck jika volume transaksi benar-benar ekstrem (puluhan juta per detik).
		        
		    -   **Penyimpanan Pesan Sementara:** Pesan biasanya dihapus setelah dikonsumsi dan dikonfirmasi. Tidak dirancang untuk menyimpan data streaming dalam jangka panjang untuk diproses ulang. 
		    
		**1. Rekomendasi Utama (Primary Choice): RabbitMQ**

		-   "**Untuk solusi utama, saya merekomendasikan RabbitMQ.** Berdasarkan requirement, prioritasnya adalah reliable task processing untuk checkin/checkout, flexible routing, dan operasional yang relatif sederhana. RabbitMQ sangat cocok untuk pola ini dengan fitur seperti..."
		    
		-   _(Jelaskan poin-poin yang sudah kita diskusikan: queue-based, ACK/NACK, DLQ, exchange types, dll)._
		    

		**2. Opsi Alternatif & Skenario Penggunaannya (Alternative Choice): Apache Kafka**

		-   "**Sebagai opsi alternatif yang powerful, kita dapat mempertimbangkan Apache Kafka.** Kafka akan unggul jika kita mengantisipasi volume data yang sangat masif dan membutuhkan kemampuan _event streaming_ dan _reprocessing_ data."
		    
		-   _(Jelaskan kekuatan Kafka: throughput, durability, stream processing)._
		    
		-   "**Namun, perlu dipertimbangkan bahwa** operasional Kafka lebih kompleks dan mungkin _overkill_ jika kebutuhan utamanya adalah task distribution yang sederhana."
		    

		**3. Kesimpulan dan Hybrid Approach (Opsional)**

		-   "**Kesimpulannya, RabbitMQ adalah pilihan yang tepat untuk inti bisnis proses checkin/checkout.**"
		    
		-   "**Namun, dalam skala yang sangat besar, pola hybrid juga mungkin,** di mana RabbitMQ menangani proses inti transaksi, sementara Kafka digunakan untuk menangani aliran _event_ untuk keperluan _analytics, monitoring, dan audit_ dalam skala besar." 



2.  ### **Data Layer :**
    
	- ### **PostgreSQL: Robust, Extensible, and Enterprise-Grade Relational Database**

		-   **ACID-Compliant for Financial Transactions**: **Menjamin Atomicity, Consistency, Isolation, and Durability untuk setiap transaksi.** Ini adalah fitur **paling kritis** untuk operasi seperti deduct saldo saat checkout dan top-up, yang harus absolut akurat dan bebas dari race condition.
    
		-   **Data Integrity & Constraints**: **Mendukung foreign keys, unique constraints, dan check constraints.** Memastikan integritas relational data terjaga (e.g., setiap `transaction` harus punya `card_id` yang valid, `balance` tidak boleh negatif). **Mencegah data corrupt dari level yang paling dasar.**
    
		-   **Rich Data Types for Flexibility**: **Support untuk JSONB, Array, dan UUID.** JSONB sangat berguna untuk menyimpan data profil user atau metadata transaksi yang dinamis tanpa perlu mengubah schema tabel secara terus-menerus.
    
		-   **Scalability for Growth**: **Mendukung built-in replication (streaming) untuk high availability dan read scaling.** Serta table partitioning untuk mengelola tabel transaksi yang membesar dengan cara mempartisi berdasarkan waktu (e.g., partisi per bulan), sehingga menjaga performa query tetap optimal.
    
		-   **Spatial Capabilities (with PostGIS)**: **Dapat diperluas dengan extension PostGIS.** Ini membuka kemungkinan untuk melakukan query geospatial, seperti mencari terminal terdekat atau menganalisis rute perjalanan penumpang di masa depan.
    
		-   **Proven Reliability & Ecosystem**: **Open-source, mature, dan memiliki ecosystem tools yang lengkap** untuk management (PgAdmin), high availability (Patroni), dan monitoring. **Pilihan yang aman dan dapat diandalkan untuk sistem bisnis kritikal 24/7.**
	- ### **Redis: Blazing-Fast, In-Memory Data Store for Low-Latency Operations**

		-   **In-Memory Performance for Critical Paths**: **Menyimpan data di RAM untuk operasi baca/tulis dalam microseconds.** Sangat cocok untuk **validasi saldo kartu dan cek status kartu langsung di gate** yang membutuhkan latensi sangat rendah (<100ms) untuk pengalaman tap-and-go yang mulus.
    
		-   **Rich Data Structures for Flexibility**: **Mendukung berbagai tipe data seperti Hashes, Sets, dan Sorted Sets.** Memungkinkan penyimpanan data terstruktur seperti **profil kartu (Hashes)** dan **daftar tarif per terminal (Sorted Sets)** secara efisien dalam satu cache layer.
    
		-   **Distributed Caching with Redis Cluster**: **Mendukung sharding dan replication untuk membangun cluster yang highly available dan scalable.** Memastikan bahwa beban dari puluhan ribu gate dapat ditangani tanpa penurunan performa, dan sistem tetap berjalan bahkan jika beberapa node cache mati.
    
		-   **Persistence for Data Safety**: **Meskipun in-memory, Redis menawarkan opsi persistence (RDB snapshots dan AOF logs) untuk mencegah kehilangan data secara total.** Sangat penting untuk menyinkronkan ulang data saldo dari database utama setelah outage tanpa harus kehilangan data transaksi terbaru.
    
		-   **Ideal Use Cases in E-Ticketing**:
    
		    -   **Hot Data Cache:** Menyimpan **saldo terkini** dan **status kartu** (active/blocked) untuk akses super cepat di gate.
        
		    -   **Fare Table Cache:** Menyimpan seluruh **matriks tarif** antar terminal yang dibaca secara konstan.
        
		    -   **Rate Limiting & Circuit Breaker:**  **Membatasi jumlah request dari sebuah gate** yang bermasalah untuk mencegah overload pada service backend dan menerapkan pola circuit breaker.
        
		    -   **Distributed Locking:**  **Mengkoordinasi akses ke resource bersama** (e.g., saat top-up saldo dari dua sumber berbeda secara bersamaan) untuk mencegah race condition.
        
	- ### **SQLite: Lightweight, Embedded, and Reliable Offline Database**

		-   **Serverless & Zero-Configuration for Simplicity**: **Beroperasi tanpa server terpisah, hanya menggunakan file .db pada sistem file gate.** Ini menghilangkan kompleksitas management server database pada setiap perangkat gate, membuat deployment, update, dan maintenance perangkat menjadi sangat sederhana dan cepat.
    
		-   **ACID-Compliant for Offline Transaction Safety**: **Menjamin Atomicity, Consistency, Isolation, and Durability untuk setiap transaksi yang terjadi di gate.** Ini adalah fitur **paling kritis** untuk operasi offline. Setiap tap kartu (check-in/check-out) akan tercatat dengan aman dan konsisten di penyimpanan lokal, bahkan selama internet mati, mencegah kehilangan data atau korupsi transaksi.
    
		-   **Unmatched Reliability in Constrained Environments**: **Dirancang untuk berjalan stabil pada perangkat dengan resource terbatas (CPU rendah, RAM kecil, dan power tiba-tiba mati).** Tidak seperti database server yang mungkin crash dalam kondisi tersebut, SQLite telah terbukti sangat robust (digunakan di Android, iOS, desktop app), menjamin gate tetap beroperasi 24/7 dalam segala kondisi.
    
		-   **Ideal for Store-and-Forward Pattern**: **Menyediakan penyimpanan sementara (local stash) yang sangat andal untuk transaksi yang belum terkirim.** Semua event tap akan langsung dicatat di tabel `transactions_pending` di SQLite. Sebuah background process kemudian akan mengirimkan data ini ke cloud secara batch ketika koneksi internet kembali tersedia.
    
		-   **Cross-Platform Portability & Wide Adoption**: **Database file-nya dapat dibaca dan ditulis oleh hampir semua platform pemrograman dan OS.** Hal ini memudahkan pengembangan firmware untuk gate dan memastikan data dapat dipulihkan (recovery) dengan mudah jika diperlukan. **Merupakan pilihan yang sudah teruji (battle-tested) untuk skenario offline-first.**
        
        
3.  ### **Infrastructure :**
	-	### **Kubernetes: The Resilient, Scalable, and Automated Orchestration Platform**

		-   **High Availability & Self-Healing for 24/7 Uptime:**  **Secara otomatis merestart, mengganti, dan menjadwalkan ulang pod yang gagal atau tidak responsif.** Ini adalah fitur **paling kritis** untuk memastikan semua service inti (sinkronisasi, top-up, API) tetap berjalan tanpa henti. Jika sebuah node server cloud mati, Kubernetes secara instan memindahkan beban kerja ke node yang sehat, memberikan ketahanan terhadap kegagalan infrastruktur.
    
		-   **Horizontal Autoscaling for Bursty Traffic Loads:**  **Dapat secara otomatis menambah atau mengurangi jumlah instance pod (replica) berdasarkan metrics CPU, memory, atau custom metrics (e.g., panjangnya antrian sinkronisasi).** Sangat cocok untuk menangani **gelombang traffic yang tiba-tiba masif**, seperti ketika ratusan gate melakukan sinkronisasi secara bersamaan setelah pemadaman internet berakhir, tanpa perlu intervensi manual.
    
		-   **Declarative Configuration & GitOps for Agile Operations:**  **Seluruh state yang diinginkan (aplikasi, konfigurasi, secret) dideklarasikan dalam file YAML/Helm yang dapat disimpan di version control (Git).** Memungkinkan **Infrastructure as Code (IaC)**, yang menyederhanakan deployment yang konsisten di berbagai environment (dev/staging/prod), rollback yang cepat dan aman jika terjadi error, dan otomatisasi CI/CD.
    
		-   **Service Discovery & Load Balancing for Resilient Microservices:**  **Setiap service mendapatkan IP dan DNS name yang stabil di dalam cluster.** Kubernetes secara otomatis mendistribusikan lalu lintas jaringan ke semua pod yang sehat di belakang sebuah service, menjamin bahwa kegagalan satu instance tidak mengganggu keseluruhan aplikasi dan komunikasi antar microservices tetap lancar.
    
		-   **Seamless Updates with Rolling & Canary Deployments:**  **Mendukung strategi deployment tanpa downtime (zero-downtime).** Memungkinkan update versi service (e.g., service pricing) dilakukan secara bertahap (**rolling update**) atau diuji coba ke sebagian kecil traffic terlebih dahulu (**canary deployment**) untuk memastikan stabilitas sebelum rollout penuh, meminimalkan risiko outage.
    
		-   **Multi-Cloud & Hybrid Cloud Flexibility:**  **Berjalan secara konsisten di berbagai environment cloud (AWS, GCP, Azure) maupun on-premise.** Ini memberikan fleksibilitas strategis untuk menghindari vendor lock-in dan membangun arsitektur hybrid yang mungkin diperlukan untuk compliance atau disaster recovery.
    
    -   ### **Istio: The Intelligent Service Mesh for Control, Security, and Visibility**

		-   **Advanced Traffic Management for Zero-Downtime Operations:**  **Mengendalikan lalu lintas jaringan antar microservices dengan presisi tinggi.** Memungkinkan strategi deployment seperti **Canary** (mengarahkan sebagian traffic ke versi baru untuk testing) dan **Blue-Green** (switch instan antara versi lama dan baru) **tanpa downtime**. Fitur **retry, timeout, dan failover** otomatis membuat sistem secara keseluruhan lebih tahan banting (resilient) terhadap kegagalan sementara di service individual.
    
		-   **Automatic Mutual TLS (mTLS) for Service-to-Service Security:**  **Mengenkripsi dan mengotentikasi semua komunikasi antar pod dalam cluster secara otomatis.** Istio yang mengelola sertifikat dan identitas setiap service, menghilangkan kompleksitas implementasi TLS dari kode aplikasi. Ini menjamin bahwa data sensitif (seperti transaksi dan saldo) diamankan dari serangan man-in-the-middle, bahkan di dalam jaringan internal.
    
		-   **Unified Observability for End-to-End Monitoring and Debugging:**  **Menyediakan metrics, logs, dan traces yang detail dan terintegrasi untuk semua komunikasi layanan.** Memberikan visibilitas penuh (**full-stack visibility**) untuk melacak sebuah permintaan dari API Gateway, melalui berbagai microservices, hingga ke database. Ini sangat penting untuk **menganalisis bottleneck performa**, mendebug issue yang kompleks, dan memahami dependencies antar service.
    
		-   **Resilience Policies with Circuit Breaking and Rate Limiting:**  **Menerapkan kebijakan ketahanan seperti circuit breaker dan pembatasan kuota (rate limiting) pada tingkat jaringan.** Mencegah kegagalan satu service menyebar (cascading failure) ke seluruh sistem dengan cara membuka circuit jika sebuah service overload. Dapat juga membatasi jumlah panggilan ke service kritikal (seperti Payment Service) untuk melindunginya dari beban berlebih (overload) dan serangan denial-of-service (DoS).
		
	- ### **Hashicorp Vault: Centralized Secrets Management and Data Protection**

		-   **Centralized & Secure Secrets Management:**  **Berfungsi sebagai sumber kebenaran tunggal yang terenkripsi untuk semua secret (credentials database, API keys, certificate TLS).** Menghilangkan praktik berbahaya seperti menyimpan secret dalam file konfigurasi atau environment variable dalam bentuk plaintext, sehingga secara drastis mengurangi risiko eksposisi secret yang dapat menyebabkan pelanggaran data.
    
		-   **Dynamic Secrets for Least-Privilege Access:**  **Dapat menghasilkan credential database (e.g., untuk PostgreSQL) yang bersifat dinamis dan berumur pendek (ephemeral) secara on-demand.** Ini memastikan bahwa setiap service atau pod hanya mendapatkan akses yang diperlukan untuk jangka waktu yang sangat terbatas (misalnya 1 jam). **Risiko keamanan diminimalkan secara signifikan** karena credential yang bocor menjadi tidak valid dalam waktu singkat.
    
		-   **Encryption as a Service for Application Data:**  **Menyediakan API sederhana bagi aplikasi untuk mengenkripsi dan mendekripsi data sensitif tanpa harus mengelola kunci enkripsi mereka sendiri.** Aplikasi hanya perlu mengirim data plaintext ke Vault untuk dienkripsi dan menyimpan ciphertext-nya. Ini menyederhanakan manajemen kunci (key rotation, key revocation) dan meningkatkan keamanan data aplikasi.
    
		-   **Fine-Grained, Policy-Based Access Control:**  **Akses ke setiap secret diatur oleh kebijakan (policies) yang terdefinisi dengan jelas berdasarkan prinsip least privilege.** Setiap service (melalui identity-nya) hanya diizinkan membaca secret yang secara eksplisit diperbolehkan untuknya. **Mencegah lateral movement** dan membatasi dampak jika sebuah pod disusupi.
    
		-   **Native Kubernetes & Istio Integration for Automated Workflows:**  **Terintegrasi erat dengan Kubernetes Service Accounts untuk memberikan secret secara otomatis ke dalam pod saat startup.** Pod tidak perlu mengetahui credential Vault; integrasi ini memberikan autentikasi yang aman dan transparan. Digabungkan dengan Istio mTLS, ini menciptakan zero-trust network yang komprehensif, dari autentikasi service hingga akses secret.
    
		-   **Comprehensive Audit Logging for Compliance and Forensics:**  **Mencatat setiap akses dan operasi pada vault (siapa, kapan, secret apa yang diakses, dan apa yang dilakukan).** Log audit yang tidak dapat diubah (immutable) ini sangat penting untuk memenuhi kebutuhan compliance, audit keamanan, dan investigasi forensik jika terjadi insiden keamanan.
     
5.  **Observability & Monitoring**
    
	- ### **OpenTelemetry: The Unified Standard for End-to-End Observability**

		-   **Vendor-Neutral Telemetry Data Collection:**  **Menyediakan APIs, SDKs, dan tools yang standar untuk mengumpulkan dan mengekspor telemetry data (metrics, logs, traces).** Menghilangkan lock-in terhadap vendor observability tertentu. Data dapat dikirim ke backend pilihan seperti Prometheus, Jaeger, atau Grafana Cloud, memastikan fleksibilitas dan **future-proofing** untuk infrastruktur monitoring.
    
		-   **Unified Framework for Simplified Management:**  **Menggabungkan tiga pilar observability (metrics, logs, traces) ke dalam satu framework dan pipeline data yang terpadu.** Ini menyederhanakan manajemen instrumentasi, mengurangi overhead operasional, dan memberikan konsistensi data dibandingkan dengan mengelola tools yang terpisah (e.g., Prometheus untuk metrics, ELK untuk logs, Jaeger untuk traces).
    
		-   **Powerful Distributed Tracing for Root Cause Analysis:**  **Memberikan visibility yang lengkap atas jalur sebuah request yang melintasi banyak service.** Setiap transaksi tap kartu menghasilkan **trace** yang menunjukkan perjalanannya dari API Gateway, ke Auth Service, lalu ke Transaction Service, dan akhirnya ke database. Ini sangat penting untuk **mengidentifikasi latency bottleneck**, mendebug error yang kompleks, dan memahami dependencies antar service.
    
		-   **Context Propagation for Correlating Telemetry Data:**  **Secara otomatis menyebarkan context (trace ID, span ID) antar service yang saling berkomunikasi.** Memungkinkan untuk mengelompokkan dan menghubungkan semua metrics, logs, dan traces yang berasal dari satu request tertentu di seluruh boundary service. **Menghilangkan tebakan** dalam melacak alur transaksi yang terdistribusi.
    
		-   **Rich Instrumentation Ecosystem:**  **Menawarkan libraries yang luas untuk bahasa pemrograman populer (Go, Java, Python, JS) dan framework yang umum (HTTP, gRPC, database drivers).** Memungkinkan developer untuk menginstrumentasi kode mereka dengan mudah, baik secara manual maupun otomatis (auto-instrumentation), sehingga mempercepat adopsi observability yang komprehensif.
        
    -   ### **Prometheus: The Time-Series Database for Operational Metrics and Alerting**

		-   **High-Performance Time-Series Database for Metrics:**  **Menyimpan semua data metrik dalam format time-series yang efisien, dilabeli dengan dimensi yang kaya (e.g., `gate_transaction_duration_seconds{terminal="A",status="success"}`).** Format ini sangat cocok untuk data operasional seperti jumlah request, latency, error rate, dan utilisasi resource, yang merupakan inti dari memonitor sistem yang berjalan 24/7.
    
		-   **Pull-Based Model with Dynamic Service Discovery:**  **Secara aktif menarik (scrape) data metrik dari endpoint HTTP (/metrics) pada target yang dimonitor.** Terintegrasi sempurna dengan Kubernetes Service Discovery untuk secara otomatis menemukan dan mulai memonitor pod dan service baru yang di-deploy, sangat menyederhanakan manajemen dalam lingkungan yang dinamis.
    
		-   **Powerful PromQL for Deep System Analysis:**  **Menyediakan bahasa query (PromQL) yang sangat fleksibel dan ekspresif untuk menganalisis data metrik.** Memungkinkan pembuatan query kompleks seperti: "**Rata-rata latency checkout dalam 5 menit terakhir, dikelompokkan per terminal**" atau "**Rasio error rate terhadap total request per service**", memberikan insight yang mendalam untuk troubleshooting dan kapasitas planning.
    
		-   **Proactive Alerting with Alertmanager for Incident Prevention:**  **Terintegrasi dengan Alertmanager untuk mendefinisikan dan mengirimkan notifikasi berdasarkan kondisi abnormal.** Dapat mengonfigurasi alert seperti: "**Kirim peringatan ke Slack/Teams jika error rate sync service melebihi 10% selama 2 menit**" atau "**Page engineer on-call jika throughput transaksi turun drastis**", memungkinkan tim untuk bertindak **sebelum** insiden menjadi kritis.
    
		-   **Efficient and Scalable for High-Volume Systems:**  **Dirancang untuk reliabilitas dan efisiensi dalam skala besar.** Dapat menangani jutaan sampel metrik per detik dari seluruh gate dan service. Strategi federasi dan retention policy yang cerdas memastikan operasi yang smooth dan berkelanjutan bahkan untuk sistem ticketing dengan volume transaksi yang sangat tinggi.
    
		-   **Critical Use Cases for the Ticketing Ecosystem:**  **Merupakan tulang punggung untuk memonitor kesehatan dan performa sistem secara real-time.** Secara khusus digunakan untuk: **- Melacak latency proses tap-in/tap-out di setiap gerbang. - Memantau utilisasi resource (CPU, memory) pada pod service dan perangkat gate. - Mendeteksi anomaly dalam pola transaksi yang dapat mengindikasikan gangguan. - Menyediakan data untuk dashboard operasional dan bisnis.**

	- ### **Jaeger: The Distributed Tracing System for End-to-End Latency and Fault Analysis**

		-   **End-to-End Request Flow Visualization for Complex Transactions:**  **Memungkinkan visualisasi lengkap dari jalur sebuah request yang melintasi banyak service dan boundary infrastruktur.** Untuk sebuah tap kartu, trace akan menunjukkan perjalanannya mulai dari **Gate Device → API Gateway → Auth Service → Transaction Service → PostgreSQL**, memberikan gambaran holistik yang sangat berharga untuk memahami alur sistem.
    
		-   **Deep Performance Analysis with Latency Breakdown:**  **Memecah latency total sebuah transaksi menjadi latency untuk setiap hop/service individual.** Fitur ini sangat penting untuk **mengidentifikasi bottleneck performa** secara tepat, misalnya menentukan apakah latency tinggi berasal dari koneksi database yang lambat, proses bisnis di service payment, atau jaringan ke gate terminal.
    
		-   **Rapid Root Cause Analysis for Error Investigation:**  **Menunjukkan titik kegagalan yang spesifik dan penyebarannya (cascading failure) dalam sebuah request yang gagal.** Saat terjadi insiden seperti double deduction saldo, trace langsung mengarahkan engineer pada service dan operasi database yang menjadi akar masalah, **secara drastis mempercepat waktu resolusi (MTTR)**.
    
		-   **Intelligent Sampling for Scalability and Cost Efficiency:**  **Mendukung berbagai strategi sampling (e.g., probabilistic, rate-limiting) untuk hanya menyimpan sebagian trace.** Ini memungkinkan tim untuk mendapatkan insight yang representatif tanpa harus menyimpan setiap trace dari jutaan transaksi harian, mengoptimalkan biaya storage dan processing tanpa mengorbankan value observability.
    
		-   **Seamless Integration with OpenTelemetry for a Standardized Pipeline:**  **Berperan sebagai backend dan UI untuk menampilkan trace data yang dikumpulkan oleh OpenTelemetry (OTel).** Arsitektur ini memisahkan concern collection (OTel) dan visualization (Jaeger), memastikan fleksibilitas dan menghindari vendor lock-in, sekaligus menyederhanakan stack observability.
    
		-   **Critical Use Cases for Transaction Integrity and Debugging:**  **Merupakan tool utama untuk menjamin integritas transaksi dan mendebug proses bisnis yang kompleks.** Secara khusus digunakan untuk: **- Melacak dan mendebug kasus duplikasi transaksi atau selisih data. - Menganalisis performa proses sinkronisasi data dari SQLite (offline) ke PostgreSQL (online). - Memetakan dependencies antar service untuk perencanaan kapasitas dan impact analysis.**

	- ### **Grafana Loki: Efficient and Lightweight Centralized Log Aggregation**

		-   **Label-Based Indexing for Extreme Storage Efficiency:**  **Mengindeks log hanya berdasarkan metadata (labels seperti `service_name`, `terminal_id`, `level`) dan menyimpan konten log itu sendiri dalam bentuk terkompresi tanpa index.** Pendekatan ini **secara drastis mengurangi biaya storage dan kebutuhan komputasi** dibandingkan solusi seperti Elasticsearch yang mengindeks seluruh teks, membuatnya ideal untuk agregasi log dalam skala besar dari ribuan gate.
    
		-   **Tight Integration with Grafana for Unified Observability:**  **Terintegrasi native dengan Grafana, menyediakan tampilan terpadu untuk metrics (dari Prometheus), traces (dari Tempo/Jaeger), dan logs.** Memungkinkan **correlated debugging** yang powerful: misalnya, mengklik titik latency tinggi pada graph metrics lalu langsung melihat log error yang terkait dari service tersebut pada timeframe yang sama, **tanpa perlu berpindah tools**.
    
		-   **Kubernetes-Native Log Collection with Promtail/FluentBit:**  **Memanfaatkan agen seperti Promtail atau FluentBit yang berjalan sebagai DaemonSet untuk secara otomatis mengumpulkan log dari semua pod dan node.** Konfigurasi yang sederhana untuk mengekstrak labels langsung dari metadata Kubernetes, memastikan semua log dari seluruh komponen sistem terkirim ke Loki tanpa effort manual yang signifikan.
    
		-   **Powerful LogQL for Advanced Log Querying and Aggregation:**  **Menggunakan bahasa query LogQL yang sintaksnya mirip dengan PromQL.** Memungkinkan tidak hanya filtering log dasar, tetapi juga **agregasi, pembuatan metric dari log (metric queries), dan perhitungan statistik** (e.g., hitung rate error per menit untuk service tertentu), menyatukan analisis log dan metrics dalam workflow yang konsisten.
    
		-   **Operational Simplicity and Scalability:**  **Dirancang untuk mudah dioperasikan dan diskalakan.** Dapat dimulai dengan deployment monolithic yang sederhana untuk environment development, dan diskalakan ke arsitektur distributed (microservices mode) yang highly available untuk menangani volume log yang sangat besar di production, semuanya dengan operational overhead yang rendah.
		
	- ### **Grafana: The Unified Visualization and Operational Dashboard Platform**

		-   **Single Pane of Glass for Correlated Observability:**  **Menyatukan data metrics (dari Prometheus), traces (dari Jaeger/Tempo), dan logs (dari Loki) dalam satu dashboard dan tampilan yang terintegrasi.** Ini merupakan **pengubah permainan (game-changer)** untuk root cause analysis, memungkinkan engineer untuk berpindah dari graph latency yang anomaly, langsung melihat trace yang error, dan kemudian melihat log detail dari service terkait—semuanya tanpa berpindah tool.
    
		-   **Powerful, Customizable, and Interactive Visualizations:**  **Menyediakan library panel visualisasi yang kaya (graphs, stat, gauges, heatmaps) yang sangat dapat disesuaikan.** Memungkinkan pembuatan dashboard operasional yang jelas dan informatif, seperti **Live Map Status Terminal**, **Transaction Throughput & Error Rate**, dan **Gateway Health Dashboard**, yang dapat disesuaikan untuk kebutuhan tim dev, ops, dan bisnis.
    
		-   **Centralized Alerting and Notification Hub:**  **Berfungsi sebagai pusat kontrol untuk mendefinisikan dan mengelola alert rules berdasarkan data dari semua sumber terhubung (Prometheus, Loki, etc.).** Dapat mengonfigurasi alert seperti "**Alert jika throughput transaksi turun di bawah X per menit**" dan mengirim notifikasi ke berbagai channel (Slack, PagerDuty, Email), memastikan tim selalu proaktif terhadap potensi insiden.
    
		-   **Data Source Agnostic for a Unified View:**  **Mendukung koneksi ke puluhan data source yang berbeda, mulai dari time-series database, log aggregators, SQL databases, hingga cloud monitoring tools.** Ini memungkinkan pembuatan dashboard komprehensif yang tidak hanya menampilkan data teknis tetapi juga data bisnis (e.g., **jumlah penumpang vs. pendapatan**), memberikan insight lengkap tentang kesehatan sistem.
    
		-   **Collaboration and Governance Features:**  **Mendukung fitur kolaborasi seperti shared dashboards, folder organization, dan Role-Based Access Control (RBAC).** Memungkinkan tim yang berbeda (Network, Backend, Hardware) untuk membuat dan berbagi dashboard mereka sendiri, sementara admin dapat mengontrol aksesnya, mempromosikan kolaborasi yang teratur dan aman.

        
