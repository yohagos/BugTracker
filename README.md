# BugTracker

Dieses Projekt wurde aufgesetzt, damit ich meine Go Kenntnisse und Erfahrungen erweiteren bzw. vertiefen kann. Hierzu entschloss ich mich MongoDB als Datenbank zu verwenden, um somit eine weitere NoSQL basierende Datenbank meinen Erfahrungen und Kenntnissen hinzu zu fügen. Desweiteren möchte ich ebenfalls meine Docker Kenntnisse und Erfahrungen erweitern und möchte diese Applikation anschließend auch als Docker Image bereitstellen.

Das Projekt soll zum nächstmöglichen Zeitpunkt, sobald nach meinem Ermessen ein funktionstüchtiger Stand (Prototype) erreicht wurde, ein Docker Image gebaut und zum testen bereitgestellt werden.

Die folgenden Bibliotheken bzw. GitHub Repositorien wurden zur Erstellung bezogen und im Projekt genutzt:

- github.com/gorilla/mux:
Gorilla Mux (als Multiplexor / Server) stellt alle benötigten Funktionalitäten bereit, um die Client Requests zu verarbeiten und die Server Response bereitzustellen.

- github.com/gorilla/sessions:
Gorilla Sessions (zur Session-Verwaltung) wird verwendet um die User Sessions als Cookies zu speichern und somit einen grundlegenden Zugriffsschutz bereitzustellen.

- mongodb:
Diverse Bibliotheken die von MongoDB bereitgestellt werden, so dass eine Verbindung zur Datenbank erzeugt werden kann und CRUD Operations durchgeführt werden können.


Im Laufe des Developments möchte ich ebenfalls die Funktionen und Möglichkeiten von GitHub Actions kennenlernen und nutzen, um meinen CI/CD Prozess zu beschleunigen.

Die Applikation beinhaltet folgende Funktionalitäten (gegenwärtiger Stand):
- Dynamische Templates (*.html):
  HTML5 Templates die mit Go  Elementen erweitert werden.
- Webserver:
  Einen Mux Router der sämtliche Requests vom Client verarbeitet und die Response an den Client zurück schickt.
- Datenbank Controller:
  Eine Connection zur MongoDB (Docker Container) die Functions für CRUD Operations bereitstellt.
- Models:
  (Gegenwärtig nur User) Ein Package für alle Data Models, wie z.B. Users, so dass für jedes Objekt (structs) ihre speziellen Funktionen bereitstehen.
- Utilities:
  Funktionalitäten die unregelmäßig verwendet werden bzw. nur an gewissen stellen, wie z.B. alle selbst definierten Errors (Fehlermeldungen, die für die Applikation definiert wurden).

Hinweise:
- Ich versuche alle relevanten Regeln von "Clean Code" zu befolgen. Der VS Code 'Go Plugin' gibt jedoch vor, dass alle global verfügbaren Objekte (Structs, Variablen, etc.) einen Kommentar erhalten, ansonsten werden diese als Warnings erfasst. Somit werden so einige "nicht relevante" Kommentare im Code vorkommen.

- Meine berufliche Tätigkeit nimmt mehr Zeit in Anspruch als erwartet, wodurch die Implementierung dieser Applikation nur langsam voran geht.
