# Berechnung und Anwendung von Quersummen

Die Aufgaben in diesem Package befassen sich mit der Berechnung von Quersummen.

## Hintergrund

Die Ziffern einer Zahl aufzusummieren, ist die Grundlage für verschiedene
Verfahren zur Analyse von Zahlen-Eigenschaften und zur Fehlererkennung.

Die einfachsten Anwendungen sind schon aus der Schule bekannt, z.B. die
Prüfung, ob eine Zahl durch 3 oder 9 teilbar ist, indem man ihre Ziffern
aufsummiert und prüft, ob die resultierende Quersumme durch 3 bzw. 9 teilbar ist.

Komplexere Varianten finden sich z.B. in der Berechnung von Prüfziffern
für Zahlenfolgen, die eingegeben oder eingescannt werden, wie etwa
Barcodes oder Kreditkartennummern.
Die Idee hierbei ist, dass sich durch falsch eingegebene Ziffern z.B.
die Quersumme ändert, und so Fehler erkannt werden können.

## Aufgaben-Überblick und empfohlene Reihenfolge

1. In `digits.go` geht es darum, aus einer Zahl die einzelnen Ziffern zu extrahieren
   und als Liste zurückzugeben. Diese Funktion ist eine Grundlage für die weiteren Aufgaben.
2. In `digit_sum.go` sollen Sie eine Funktion implementieren, die die Quersumme
   einer Zahl berechnet, also die Summe ihrer Ziffern.
3. In `divisibility.go` implementieren Sie eine Funktion, die prüft,
   ob eine Zahl durch 3 teilbar ist, indem sie die Quersumme verwendet.
4. In `ean_checksum.go` schließlich sollen Sie die Prüfziffer
   für eine Liste von Ziffern nach dem EAN-13-Standard berechnen und prüfen.
