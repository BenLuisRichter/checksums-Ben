# Berechnung von Paritätsbits

Bei den Aufgaben in diesem Package geht es darum, Paritätsbits zu berechnen und zu setzen.

## Hintergrund

Paritätsbits sind eine einfache Methode zur Fehlererkennung in der Datenübertragung.
Bei der Übertragung von Daten in Binärform (also als Nullen und Einsen) kann es vorkommen,
dass einzelne Bits fehlerhaft übertragen werden, also von 0 zu 1 oder umgekehrt kippen.
Eine Methode, solche Fehler zu erkennen, ist die Verwendung von Paritätsbits.
Dabei wird zu einer Gruppe von Bits ein zusätzliches Bit hinzugefügt,
so dass die Gesamtzahl der Einsen in der Gruppe entweder gerade (gerade Parität)
oder ungerade (ungerade Parität) ist.

## Aufgaben-Überblick

In den Dateien `set_parity_bit.go` und `check_parity_bit.go` finden Sie Funktionen,
die Paritätsbits setzen und überprüfen. Diese sind also das Ziel der Aufgaben in diesem
Package. In `count_binary_ones.go` und finden Sie eine Hilfsfunktion zur Zählung der Einsen
in der Binärdarstellung einer Zahl. In `binary_string.go` gibt es eine Funktion,
die eine Zahl in ihre Binärdarstellung als String umwandelt. Diese ist nicht direkt
Teil der Paritätsbit-Berechnung, kann aber hilfreich sein, um sich den anderen Funktionen
zu nähern und deren Funktionsweise zu verstehen.

## Empfohlene Reihenfolge

1. `binary_string.go`: Beginnen Sie mit dieser Datei, um ein Verständnis dafür zu
   entwickeln, wie Zahlen in Binärform dargestellt werden.
   Wenn Ihnen das Konzept der Binärdarstellung bereits vertraut ist, können Sie
   diese Aufgabe auch überspringen.
2. `count_binary_ones.go`: Dies ist die eigentlich "Logik" hinter der Paritätsbit-Berechnung.
   Diese Funktion sollten Sie implementieren und dann in den weiteren Aufgaben verwenden.
3. `check_parity_bit.go`: Implementieren Sie diese Funktion, um zu überprüfen, ob
   das Paritätsbit korrekt gesetzt ist.
4. `set_parity_bit.go`: Implementieren Sie diese Funktion, um ein Paritätsbit zu setzen.
