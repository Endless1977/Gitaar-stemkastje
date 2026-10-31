# Gitaar-stemkastje

Een gitaartuner die je van de ene stemming naar de andere brengt. Kies waar je gitaar nu op staat en waar je naartoe wilt (bijvoorbeeld standaard → open G). Per snaar zie je hoeveel halve tonen hij strakker of losser moet, en de meter laat via de microfoon zien wanneer hij klopt.

**Spelen:** https://endless1977.github.io/Gitaar-stemkastje/

## Stemmingen

Standaard, Open G, Open D, Open E, Open A, Open C, Open D mineur, Drop D, Drop C, DADGAD, halve toon lager en hele toon lager.

## Gebruik

1. Kies je huidige stemming en je doelstemming.
2. Klik **Start microfoon** en sta de microfoon toe.
3. Tokkel één snaar tegelijk. De tuner kiest zelf de dichtstbijzijnde snaar; klik op een snaar om hem vast te zetten.
4. Met ▶ hoor je de doeltoon, handig om op gehoor te stemmen.

Een waarschuwing verschijnt als een snaar 3 halve tonen of meer omhoog moet.

## Windows-programma

`main.go` is een kleine launcher die de pagina lokaal serveert en in je browser opent. Bouwen:

```
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -H windowsgui" -o Stemkastje.exe .
```

Een kant-en-klare `Stemkastje.exe` staat bij de Releases.
