# Solar Energy API

Questa API REST fornisce l’accesso ai dati di **produzione** e **consumo energetico**
memorizzati in **InfluxDB**, provenienti da un impianto fotovoltaico.

L’obiettivo principale dell’API è supportare la **visualizzazione grafica dei dati**
(grafici giornalieri, confronti tra giorni, analisi su intervalli di tempo).

---

## 📊 Tipologie di interrogazione supportate

Attualmente l’API consente di interrogare i dati secondo le seguenti modalità:

- **Giorno corrente**
- **Giorno precedente**
- **Intervallo di date personalizzato**, con aggregazione giornaliera

Tutti gli intervalli giornalieri sono definiti secondo la seguente regola:

> ⏱️ Ogni giorno è considerato nell’intervallo  
> **00:00 → 24:00 (UTC)**  
> in modo da avere grafici confrontabili e con la stessa scala temporale.

---

## 🌍 Timezone

- Tutti i calcoli temporali vengono effettuati **nel backend**
- Le date sono espresse in **UTC**
- InfluxDB riceve sempre timestamp in formato **RFC3339**

Questo approccio garantisce risultati deterministici e indipendenti dal timezone del server.

---

## 🔗 Endpoints disponibili

### `GET /energy/today`

Restituisce i dati di produzione e consumo **del giorno corrente**,  
dall’inizio della giornata (00:00 UTC) fino al momento della richiesta.


### `GET /energy/yesterday`

Restituisce i dati di produzione e consumo **del giorno precedente**,  
nell’intervallo completo:
```
00:00 → 24:00 (UTC)
```

### `GET /energy/daily`

Restituisce i dati di produzione e consumo per un **intervallo di date personalizzato**,
raggruppati giorno per giorno.


**Query parameters:**

| Parametro | Descrizione |
|---------|-------------|
| `from` | Data di inizio (`YYYY-MM-DD`) |
| `to` | Data di fine (`YYYY-MM-DD`) |
| `window` | Finestra di aggregazione (es. `5m`, `10m`) |

**Esempio:**
```
GET /energy/daily?from=2026-04-01&to=2026-04-07&window=10m
```

## 🧱 Architettura (overview)

L’applicazione è strutturata in layer ben separati:

- **API layer**: gestione HTTP e routing (Gin)
- **Service layer**: logica applicativa e calcolo degli intervalli temporali
- **Repository layer**: query verso InfluxDB
- **Database**: InfluxDB

Il database non prende decisioni temporali:  
tutti i confini di tempo sono calcolati nel backend.

---

## 🚀 Use case principali

- Dashboard di monitoraggio energetico
- Confronto tra giorni diversi
- Analisi storica della produzione e del consumo