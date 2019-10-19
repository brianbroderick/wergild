### Example Query with Full Text Search ###


```
{
  room(func: eq(roomSlug, "forwell_inn_common_room")) {
    uid
    roomName
    roomSlug
    region {
      regionName
    }
    exits {
      dest {
        roomName
        pointsOfInterest @filter(alloftext(poiName, "walnut")) {
          poiName
        }
        region {
      		regionName
    		}
      }
      name: direction
    }
    pointsOfInterest {
      poiName
    }
  }
}
```