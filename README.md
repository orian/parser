# parser
a tiny tool for parsing sql.Rows.Scan result

# Example (with gorm)
```go
rows, err := a.db.Table("entities").
    Select("entities.ID, street, local, address, lat, lng, created").
    Joins("left join geocodeds on entities.ID = geocodeds.entity_id").
    Rows()
defer rows.Close()
if err != nil {
    return nil, err
}

const cols = 7
rawResult := make([][]byte, cols)

dest := make([]interface{}, cols) // A temporary interface{} slice
for i, _ := range rawResult {
    dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
}

pars := parser.New()
p := &api.Point{}
pars.AddMany(&p.ID, &p.Street, &p.Local, &p.Address, &p.Lat, &p.Lng, &p.Created)
resp := &api.ListResp{}
for rows.Next() {
    err := rows.Scan(dest...)
    if err != nil {
        log.Printf("skipping row during scan: %s", err)
        continue
    }
    err = pars.Parse(rawResult)
    if err != nil {
        log.Printf("cannot parse: %s", err)
        continue
    }
    pt := proto.Clone(p)
    resp.Points = append(resp.Points, pt)
}
```