# README

Questa cartella contiene lo stato on-going delle attivit&agrave; di migrazione e di data preparation.

* i file .relmig sono gli export del Relational Migrator
* sono presenti degli script che eseguono diverse operazioni:
  * opem-drop-collections.sh: droppano le collection nel db target
  * opem-import.sh: esegue import dei dati esportati in precedenza 
  * opem-export.sh: esporta le collection previste
  * opem-fix-migrator-process.sh: sistema qualche cosa in termini di dati e droppa gli indici creati dal migrator

I precedenti script richiamano per le varie funzionalit&agrave; alcuni js presenti nella cartella.
In linea di principio per una installazione da zero si esegua `opem-import.sh` (lo script preventivamente _droppa_ le collection per cui in situazione iniziale potrebbe
generare degli errori).

All'interno degli script &egrave; referenziata una connection string di un server M0. Si prega di cambiarla.