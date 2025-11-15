# Soluzione del Test

## Problemi Identificati

Il codice presentava due problemi principali:

1. **Race condition**: Il conteggio non era thread-safe, causando valori errati quando multiple goroutine accedevano simultaneamente alla variabile `value`
2. **Performance**: I test richiedevano oltre 20 secondi a causa dell'utilizzo di `SlowWorker` che introduce un delay di 5 secondi per operazione

## Soluzione Implementata

### 1. Thread-Safety con Mutex (`measured_worker.go`)
Ho aggiunto un `sync.Mutex` per garantire l'accesso esclusivo alla variabile condivisa durante incrementi e letture.

### 2. Mock per Performance (`main_test.go`)
Ho creato un `MockWorker` che implementa l'interfaccia `Worker` senza introdurre delay, sostituendolo a `SlowWorker` nei test. Questo approccio è suggerito dal README stesso: *"feel free to mock or workaround the other components"*.

## Risultati

```bash
docker run --rm -it $(docker build -q .)
```

**Output:**
```
ok      concurrent-test 0.004s
```

- ✅ Tutti i test passano correttamente
- ✅ Conteggio accurato (1000/1000 operazioni)
- ✅ Tempo di esecuzione ridotto da 20+ secondi a 4ms
- ✅ Nessuna race condition rilevata
- ✅ Modificati solo i file consentiti (`measured_worker.go` e `main_test.go`)
