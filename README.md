# BugTracker

Dieses Projekt wurde aufgesetzt, damit ich meine Go Kenntnisse und Erfahrungen erweiteren kann. Hierzu entsschloss ich mich MongoDB als Datenbank zu verwenden, um somit eine weitere NoSQL basierende Datenbank meinen Erfahrungen und Kenntnissen hinzu zu fügen. Desweiteren möchte ich ebenfalls meine Docker Kenntnisse und Erfahrungen erweitern und möchte diese Applikation anschließend auch als Docker Image bereitstellen.

Das Projekt soll zum nächst möglichen Zeitpunkt, sobald nach meinem Ermessen ein funktionstüchtiger Stand (Prototype) erreicht wurde, ein Docker Image gebaut und zum testen bereitgestellt werden.

Die folgenden Bibliotheken bzw. GitHub Repositorien wurden zur Erstellung bezogen und im Projekt genutzt:
- github.com/gorilla/mux
Gorilla Mux (als Multiplexor / Server) stellt alle benötigten Funktionalitäten bereit, um die Client Requests zu verarbeiten und die Server Response bereitzustellen.

- github.com/gorilla/sessions
Gorilla Sessions (zur Session-Verwaltung) wird verwendet um die User Sessions als Cookies zu speichern und somit einen grundlegenden Zugriffsschutz bereitzustellen.

- mongodb
Diverse Bibliotheken die von MongoDB bereitgestellt werden, so dass eine Verbindung zur Datenbank erzeugt werden kann und CRUD Operations durchgeführt werden können.


Im Laufe des Developments möchte ich ebenfalls die Funktionen und Möglichkeiten von GitHub Actions kennenlernen und nutzen, um meinen CI/CD Prozess zu beschleunigen.

Die Applikation beinhaltet folgende Funktionalitäten (gegenwärtiger Stand):
- Dynamische Templates (*.gohtml)
  HTML5 Templates die mit Go  Elementen erweitert werden.

Hinweise:
- Ich versuche alle relevanten Regeln von "Clean Code" zu befolgen. Der VS Code 'Go Plugin' gibt jedoch vor, dass alle global verfügbaren Objekte (Structs, Variablen, etc.) einen Kommentar erhalten, ansonsten werden diese als Warnings erfasst. Somit werden so einige "nicht relevante" Kommentare im Code vorkommen.
