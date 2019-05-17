# Welcome to my Revel test


### Start the web server:
#### make sure you move the package to $GOPATH/src/

   revel run beers

### Go to http://localhost:9000/beers and you'll see:

    [
      {
        "ID": 0,
        "Name": "La Trappe Quadrupel Oak Aged",
        "Type": "Ale",
        "Brewery": "Bierbrouwerij De Koningshoeven"
      },
      {
        "ID": 1,
        "Name": "Cocoa Psycho",
        "Type": "Porter",
        "Brewery": "BrewDog"
      },
      {
        "ID": 2,
        "Name": "American Dream",
        "Type": "Lager",
        "Brewery": "Mikkeller"
      }
    ]

### Go to http://localhost:9000/beers/0 and you'll see:
    {
      "ID": 0,
      "Name": "La Trappe Quadrupel Oak Aged",
      "Type": "Ale",
      "Brewery": "Bierbrouwerij De Koningshoeven"
    }
    
#### Same for 1 and 2
