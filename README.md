# BugTracker

Dieses Projekt wurde aufgesetzt, damit ich meine Go Kenntnisse und Erfahrungen erweiteren bzw. vertiefen kann. Hierzu entschloss ich mich MongoDB als Datenbank zu verwenden, um somit eine weitere NoSQL basierende Datenbank meinen Erfahrungen und Kenntnissen hinzu zu fügen. Desweiteren möchte ich ebenfalls meine Docker Kenntnisse und Erfahrungen erweitern und möchte diese Applikation anschließend auch als Docker Image bereitstellen.

Diese Applikation soll es seinen Nutzern die folgenden Funktionen bereitstellen:
- Account erstellen
Ein User muss seine Daten (bspw. Name, Email, etc.) eingeben und registrieren.
- Einloggen
Durch Eingabe der Email Adresse und Passwort sich einloggen.
- Kategorien von Tickets zu definieren
Es soll ermöglicht werden kundenspezifische Kategorien (bspw. Bugs, Defects, etc.) zu erstellen. 
- Tickets zu erstellen
Tickets sollen erstellt werden können, in denen Kategorien ausgewählt werden können und zusätzliche Informationen (bspw. Ticket ID, Zuständiger, Status, etc.) hinzugefügt werden können.

Das Projekt soll zum nächstmöglichen Zeitpunkt, sobald ein funktionstüchtiger Stand (Prototype) erreicht wurde, ein Docker Image gebaut und zum testen bereitgestellt werden.

Die folgenden Bibliotheken bzw. GitHub Repositorien wurden zur Erstellung bezogen und im Projekt genutzt:

- github.com/gorilla/mux:
Gorilla Mux (als Multiplexor / Server) stellt alle benötigten Funktionalitäten bereit, um die Client Requests zu verarbeiten und die Server Response bereitzustellen.

- github.com/gorilla/sessions:
Gorilla Sessions (zur Session-Verwaltung) wird verwendet um die User Sessions als Cookies zu speichern und somit einen grundlegenden Zugriffsschutz bereitzustellen.

- go.mongodb.org/mongo-driver/*:
Diverse Bibliotheken die von MongoDB bereitgestellt werden, so dass eine Verbindung zur Datenbank erzeugt werden kann und CRUD Operations durchgeführt werden können.


Im Laufe des Developments möchte ich ebenfalls die Funktionen und Möglichkeiten von GitHub Actions kennenlernen und nutzen, um meinen CI/CD Prozess zu beschleunigen.

Die Applikation beinhaltet folgende Funktionalitäten (gegenwärtiger Stand):
- Dynamische Templates (*.gohtml):
  HTML5 Templates die mit Go  Elementen erweitert werden.
- Servers:
  Einen Mux Router der sämtliche Requests vom Client verarbeitet und die Response an den Client zurück schickt. Ein Fileserver der sämtliche Dateien (.gohtml, .css & .js) für den Webserver lädt und bereitstellt.
- Database:
  Das Package Databases enthält eine Init Function, die zum einen die Verbindung zur Datenbank herstellt als auch sämtlichen Collections erstellt. Es wurden für alle CRUD Operationen eigene Service Klassen erstellt.
- Models:
  Ein Package für alle Data Models, wie z.B. Users, so dass für jedes Objekt (structs) ihre eigenen Methods und Functions bereitstehen.
- Utilities:
  Funktionalitäten die unregelmäßig verwendet werden bzw. nur an gewissen stellen.
- AppErrors:
  Ein eigenes Package für Errors wurde erstellt, worin sämtlichen Fehlermeldungen für die Applikation definiert wurden.

Hinweise:
- Ich versuche alle relevanten Regeln von "Clean Code" zu befolgen. Der VS Code 'Go Plugin' gibt jedoch vor, dass alle global verfügbaren Objekte (Structs, Variablen, etc.) einen Kommentar erhalten, ansonsten werden diese als Warnings erkannt worden. Somit werden so einige "nicht relevante" Kommentare im Code vorkommen.

- Meine berufliche Tätigkeit nimmt mehr Zeit in Anspruch als erwartet, wodurch die Implementierung dieser Applikation nur langsam voran geht.
