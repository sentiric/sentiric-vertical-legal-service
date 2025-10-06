# 🏛️ Sentiric Vertical Legal Service - Mantık ve Akış Mimarisi

**Stratejik Rol:** Hukuk ve dava yönetimi süreçlerine odaklanan dikey iş mantığını içerir. Müşterilerin dava durumlarını, sonraki duruşma tarihlerini veya yasal belge taleplerini yönetmek için bir arayüz sağlar.

---

## 1. Temel Akış: Dava Durumu Kontrolü (CheckCaseStatus)

```mermaid
sequenceDiagram
    participant Agent as Agent Service
    participant VLS as Legal Service
    participant CMS as Harici Dava Yönetim Sistemi (API)
    
    Agent->>VLS: CheckCaseStatus(case_number, client_id)
    
    Note over VLS: 1. Kimlik ve Yetkilendirme Kontrolü
    Note over VLS: 2. Harici CMS API Çağrısı
    VLS->>CMS: GET /case/{case_number}
    CMS-->>VLS: Dava Detayları
    
    Note over VLS: 3. Sonuçların Formatlanması
    VLS-->>Agent: CheckCaseStatusResponse(status: Active, next_hearing_date: "...")
```

## 2. Hassas Veri Yönetimi

Legal Service, dava bilgileri gibi son derece hassas müşteri verilerini işler. Bu nedenle, giden bağlantılarda ve kimlik doğrulamasında en katı güvenlik standartlarını (mTLS, katmanlı yetkilendirme) uygulamalıdır.