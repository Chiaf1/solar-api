l'obbiettivo di queste api è quello di fornire i dati immagazzinati nel db influxDB.

Le tipologie di possibili "interroggazioni" che ho pensato sono:
- Il giorno corrente
- la settimana precendete divisa giorno per giorno, non so se un endpoint per ogni giorno o uno solo perla settimana o tutti e due
- la possibilità di interrogare dando un periodo di tempo con il calendario (da giorno x a giorno y) sempre visualizzati giorno per giorno

L'idea era di visualizzare i dati in formato grafico e sempre avere i dati dalle ore 00:00 aller ore 24:59 in modo da avere i grafici tutti con la stessa scala più facili da confrontare

### Endpoints:
- GET /energy/today
- GET /energy/yesterday
- GET /energy/daily?from=YYYY-MM-DD&to=YYYY-MM-DD&window=10m
